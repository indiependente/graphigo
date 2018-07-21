package graph

// V is a vertex/node of the graph represented as an integer value
type V int

// G is a graph represented as an adjacency list of nodes
type G map[V][]V

// E is an edge between two vertexes u and v
type E [2]V

// NewNode returns a new node
func NewNode(i int) V {
	return V(i)
}

// Value returns the integer value of the node
func (v V) Value() int {
	return int(v)
}

// NewEdge returns an edge between the vertexes u and v
func NewEdge(u, v V) E {
	return E([2]V{u, v})
}

// NewGraph returns a new empty graph of capacity n
func NewGraph(n int) G {
	return make(G, n)
}

// Has returns true if the graph has that vertex otherwise false
func (g G) Has(v V) bool {
	_, ok := g[v]
	return ok
}

// Node return the node with given key i, or nil if not present
func (g G) Node(i int) *V {
	if !g.Has(V(i)) {
		return nil
	}
	for _, v := range g.Nodes() {
		if i == v.Value() {
			return &v
		}
	}
	return nil
}

// AddEdge adds an endge to the graph between the nodes u and v if not already present:
// returns true if added or false if already present
func (g G) AddEdge(u, v V) bool {
	if g.AreNeighbours(u, v) {
		return false
	}
	g[u] = append(g[u], v)
	return true
}

// AreNeighbours returns true if the u and v are neighbours in the graph
func (g G) AreNeighbours(u, v V) bool {
	uE, ok := g[u]
	if !ok {
		return false
	}
	vE, ok := g[v]
	if !ok {
		return false
	}
	if in(v, uE) && in(u, vE) {
		return true
	}

	return false
}

// Add adds a node to the graph
func (g G) Add(v V, e []V) bool {
	if g.Has(v) {
		return false
	}
	g[v] = e
	return true
}

func in(x V, array []V) bool {
	for _, i := range array {
		if i == x {
			return true
		}
	}
	return false
}

// Nodes returns a slice containigs all the graph's nodes
func (g G) Nodes() []V {
	nodes := make([]V, 0, len(g))
	for n := range g {
		nodes = append(nodes, n)
	}
	return nodes
}

// Edges returns a slice containing all the graph's edges
func (g G) Edges() []E {
	edges := make([]E, 0, len(g))
	for u := range g {
		edges = append(edges, g.EdgeList(u)...)
	}
	return edges
}

// EdgeList returns a slice containing all the edges hitting node u
func (g G) EdgeList(u V) []E {
	edges := make([]E, 0, len(g[u]))
	for _, v := range g[u] {
		edges = append(edges, NewEdge(u, v))
	}
	return edges
}
