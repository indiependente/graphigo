package graph

//V is a vertex/node of the graph represented as an integer value
type V int

//G is a an adjacency list of nodes
type G map[V][]V

//NewNode returns a new node
func NewNode(i int) V {
	return V(i)
}

//Value returns the integer value of the node
func (v V) Value() int {
	return int(v)
}

//NewGraph returns a new empty graph of capacity n
func NewGraph(n int) G {
	return make(G, n)
}

//Has returns true if the graph has that vertex otherwise false
func (g G) Has(v V) bool {
	_, ok := g[v]
	return ok
}

//AddEdge adds an endge to the graph between the nodes u and v if not already present:
//returns true if added or false if already present
func (g G) AddEdge(u, v V) bool {
	if g.AreNeighbours(u, v) {
		return false
	}
	g[u] = append(g[u], v)
	g[v] = append(g[v], u)
	return true
}

//AreNeighbours returns true if the u and v are neighbours in the graph
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

//Add adds a node to the graph
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
