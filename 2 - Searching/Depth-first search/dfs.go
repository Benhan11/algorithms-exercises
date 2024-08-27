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

	// Expected: 1 2 0 3 5 6 7 8 9 4 15 16 17 18 19 10 11 12 13 14
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
	traverseRecursive(vertices, discovered, graph.source)
}
func traverseRecursive(vertices map[int][]edge, discovered map[int]bool, source int) {
	discovered[source] = true
	fmt.Print(source, " ")

	for _, edge := range vertices[source] {
		nextVertex := edge.a
		if nextVertex == source {
			nextVertex = edge.b
		}

		if !discovered[nextVertex] {
			traverseRecursive(vertices, discovered, nextVertex)
		}
	}
}
