package graph_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/indiependente/graphigo/graph"
)

func TestGraph(t *testing.T) {
	const MAXNODES = 5
	const NODES = 2
	g := graph.NewGraph(MAXNODES)

	for i := 0; i < NODES; i++ {
		g.Add(graph.NewNode(i), nil)
	}

	if len(g) != NODES {
		t.Errorf("Graph should contain %d nodes, instead contains %d", NODES, len(g))
	}
}

func TestNeighbours(t *testing.T) {
	const MAXNODES = 5
	g := graph.CompleteGraph(MAXNODES)
	var i, j int
	rand.Seed(time.Now().UnixNano())
	for true {
		i = rand.Intn(MAXNODES - 1)
		j = rand.Intn(MAXNODES - 1)
		if i != j {
			break
		}
	}
	if !g.AreNeighbours(*g.Node(i), *g.Node(j)) {
		t.Errorf("The two vertexes %+v and %+v should be linked by an edge", *g.Node(i), *g.Node(j))
	}
}

func TestRandomGraph(t *testing.T) {
	const NODES = 6
	g1 := graph.RandomGraph(NODES, 0.8)
	g2 := graph.RandomGraph(NODES, 0.1)

	len_g1 := len(g1.Edges())
	len_g2 := len(g2.Edges())
	if len_g1 < len_g2 {
		t.Errorf("Graph #1 should have more edges than graph #2")
	}
}
