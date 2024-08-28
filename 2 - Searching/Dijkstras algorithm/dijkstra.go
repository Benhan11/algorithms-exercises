package main

import (
	"fmt"
	"math"
)

type edge struct {
	a, b, weight int
}

func main() {
	var edges = []edge{
		{0, 1, 3}, {0, 2, 1}, {1, 3, 4}, {1, 4, 1},
		{2, 5, 7}, {2, 6, 3}, {3, 7, 2}, {4, 7, 5},
		{5, 8, 2}, {6, 8, 3}, {7, 9, 1}, {8, 9, 4},
		{3, 5, 6}, {4, 6, 3}, {5, 9, 7},
	}

	fmt.Println("\n-- Traversal --")
	traverse(0, edges)

	fmt.Println("\n-- Find shortest path --")
	ui(edges)
}

func makeVertices(edges []edge) map[int][]edge {
	// Vertex represented as a key-value pair of int-[]edge
	vertices := make(map[int][]edge)

	// Assign vertex-edge pairs
	for _, edge := range edges {
		vertices[edge.a] = append(vertices[edge.a], edge)
		vertices[edge.b] = append(vertices[edge.b], edge)
	}

	return vertices
}

func traverse(src int, edges []edge) {
	dijkstra(src, -1, makeVertices(edges))
}

func findShortestPath(src, dst int, edges []edge) {
	dijkstra(src, dst, makeVertices(edges))
}

func dijkstra(src, dst int, vertices map[int][]edge) {
	dist := make(map[int]int)
	prev := make(map[int]int)
	queue := make(map[int]bool)

	for vertex := range vertices {
		dist[vertex] = math.MaxInt32
		prev[vertex] = -1
		queue[vertex] = true
	}

	dist[src] = 0

	for len(queue) > 0 {
		vertex := minDist(queue, dist)
		delete(queue, vertex) // O(1) underlying hash map

		// If there is a specified destionation vertex check if we've reached it
		if dst != -1 && vertex == dst {
			printShortest(prev, vertex, src)
			return
		}

		for _, edge := range vertices[vertex] {
			nextVertex := edge.a
			if nextVertex == vertex {
				nextVertex = edge.b
			}

			altDist := dist[vertex] + edge.weight
			if altDist < dist[nextVertex] {
				dist[nextVertex] = altDist
				prev[nextVertex] = vertex
			}
		}
	}

	printResults(dist, prev, src)
}

func minDist(queue map[int]bool, dist map[int]int) int {
	minVertex := -1
	minDist := math.MaxInt32

	for vertex := range queue {
		if dist[vertex] < minDist {
			minVertex = vertex
			minDist = dist[vertex]
		}
	}
	return minVertex
}

func ui(edges []edge) {
	var src, dst int
	fmt.Print("From: ")
	fmt.Scan(&src)
	fmt.Print("To:   ")
	fmt.Scan(&dst)

	findShortestPath(src, dst, edges)
	ui(edges)
}

func printResults(dist, prev map[int]int, src int) {
	fmt.Println("Distances (a -> b: distance):")
	for i, d := range dist {
		if i != src {
			fmt.Print(src, " -> ", i, ": ", d, "\n")
		}
	}

	fmt.Print("\nShortest paths from ", src, ":\n")
	for v := range prev {
		if v != src {
			fmt.Print("Vertex ", v, ": ")
			printPath(prev, v, src)
			fmt.Println(v)
		}
	}
}

func printShortest(prev map[int]int, dst, src int) {
	fmt.Print("\n{ ")
	printPath(prev, dst, src)
	fmt.Print(dst, " }\n\n")
}

func printPath(prev map[int]int, vertex, src int) {
	if vertex == src {
		return
	}
	printPath(prev, prev[vertex], src)
	fmt.Print(prev[vertex], " -> ")
}
