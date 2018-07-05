package graph_test

import (
	"testing"

	"github.com/indiependente/graphigo/graph"
)

func TestGraph(t *testing.T) {
	g := graph.NewGraph(5)
	u := graph.NewNode(0)
	v := graph.NewNode(1)
	g.Add(u, nil)
	g.Add(v, nil)
	g.AddEdge(u, v)

	if !g.AreNeighbours(u, v) {
		t.Errorf("The two vertexes %+v and %+v should be linked by an edge", u, v)
	}
}
