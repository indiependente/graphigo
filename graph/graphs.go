package graph

import (
	"math/rand"
	"time"
)

// RandomGraph returns a new graph with n nodes and edges with probability p of being created.
func RandomGraph(n int, p float64) G {
	// create graph
	graph := NewGraph(n)
	// create n nodes
	for i := 0; i < n; i++ {
		graph.Add(NewNode(i), nil)
	}
	// initialize random seed
	rand.Seed(time.Now().UnixNano())
	// create edges
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			r := rand.Float64()
			if r <= p {
				graph.AddEdge(*graph.Node(i), *graph.Node(j))
				graph.AddEdge(*graph.Node(j), *graph.Node(i))
			}
		}
	}
	return graph
}

// CompleteGraph returns a graph with n nodes where each one is connected to each other
func CompleteGraph(n int) G {
	graph := NewGraph(n)

	for i := 0; i < n; i++ {
		graph.Add(NewNode(i), nil)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			graph.AddEdge(*graph.Node(i), *graph.Node(j))
			graph.AddEdge(*graph.Node(j), *graph.Node(i))
		}
	}
	return graph
}
