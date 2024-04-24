package helpers

import "fmt"

type Graph struct {
	Vertices map[int]*Vertex
}

type Vertex struct {
	//add stuff as required
	Key   int //to know key easily given a vertex
	Edges []*Edge
}

type Edge struct {
	Weight rune
	Vertex *Vertex
}

func (g *Graph) AddVertex(key int) {
	g.Vertices[key] = &Vertex{key, []*Edge{}}
}

func (g *Graph) AddEdge(from, to int, weight rune) error {

	if _, ok := g.Vertices[from]; !ok {
		g.AddVertex(from)
		// return fmt.Errorf("source vertex not found")
	}
	if _, ok := g.Vertices[to]; !ok {
		g.AddVertex(to)
		// return fmt.Errorf("destination vertex not found")
	}

	g.Vertices[from].Edges = append(g.Vertices[from].Edges, &Edge{weight, g.Vertices[to]})

	return nil
}

func (g *Graph) GetNeighbours(key int) ([]*Vertex, error) {

	neighbours := []*Vertex{}
	vertex, ok := g.Vertices[key]
	if ok {
		for _, edge := range vertex.Edges {
			neighbours = append(neighbours, edge.Vertex)
		}
	} else {
		return nil, fmt.Errorf("vertex doesn't exist")
	}
	return neighbours, nil
}

// constructors from adjacency list or matrices
// can replace func(*Graph) with a type
func NewGraph(graphopts ...func(*Graph)) *Graph {
	// initiate an empty graph
	g := &Graph{
		Vertices: make(map[int]*Vertex),
	}

	for _, opt := range graphopts {
		opt(g)
	}

	return g
}

func WithAdjacencyList(directed bool) func(*Graph) {
	return func(g *Graph) {
	}
}

func GiveAllEdgesWeight(weight rune) func(*Graph) {
	return func(g *Graph) {
		// unoptimised takes 2e time
		for _, vertex := range g.Vertices {
			for _, edge := range vertex.Edges {
				edge.Weight = weight
			}
		}
	}
}

// implement constructor from matrices
