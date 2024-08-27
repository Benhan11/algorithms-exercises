package main

import (
	"fmt"
)

type edge struct {
	a, b int
}

type graph struct {
	edges  []edge
	source int
}

func main() {
	var graph = graph{
		edges: []edge{
			{1, 2}, {1, 0}, {0, 2}, {2, 3}, {2, 4},
			{3, 5}, {5, 6}, {6, 7}, {7, 8}, {8, 9},
			{1, 10}, {10, 11}, {11, 12}, {12, 13}, {13, 14},
			{4, 15}, {15, 16}, {16, 17}, {17, 18}, {18, 19},
		},
		source: 1,
	}

	// Expected: 1 2 0 10 3 4 11 5 15 12 6 16 13 7 17 14 8 18 9 19
	fmt.Print("\nTraversal: ")
	traverse(graph)
}

func traverse(graph graph) {
	// Vertex represented as a key-value pair of int-[]edge
	vertices := make(map[int][]edge)

	// Assign vertex-edge pairs
	for _, edge := range graph.edges {
		vertices[edge.a] = append(vertices[edge.a], edge)
		vertices[edge.b] = append(vertices[edge.b], edge)
	}

	discovered := make(map[int]bool)
	traverseBFS(vertices, discovered, graph.source)
}
func traverseBFS(vertices map[int][]edge, discovered map[int]bool, source int) {
	queue := make([]int, 0)

	discovered[source] = true
	queue = append(queue, source)

	for len(queue) > 0 {
		// Dequeue
		vertex := queue[0]
		queue = queue[1:]

		fmt.Print(vertex, " ")

		// Iterate over the edges of the current vertex, not the original source
		for _, edge := range vertices[vertex] {
			nextVertex := edge.a
			if nextVertex == vertex {
				nextVertex = edge.b
			}

			if !discovered[nextVertex] {
				discovered[nextVertex] = true
				queue = append(queue, nextVertex)
			}
		}
	}
}
