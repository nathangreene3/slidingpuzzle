package slidingpuzzle

import "sort"

// Vertex ...
type Vertex struct {
	pzl   *Puzzle
	nbors []*Vertex
}

// NewGraph ...
func NewGraph(pzl *Puzzle) *Vertex {
	return &Vertex{pzl: pzl, nbors: make([]*Vertex, 0, 4)}
}

func (g *Vertex) generateNeighbors() {
	opts := g.pzl.options()
	for i := range opts {
		newPzl := g.pzl.copy()
		newPzl.slide(opts[i])
		g.nbors = append(g.nbors, NewGraph(newPzl))
	}

	sort.Slice(g.nbors, func(i, j int) bool { return g.nbors[i].compareTo(g.nbors[j]) < 0 })
}

func (g *Vertex) search(pzl *Puzzle, cost float64) (float64, []*Vertex) {
	// A* minimizes a function f given an accumulating cost and heuristic (entropy).
	// f(n) = (cost+1) + n.entropy, where n is a neighbor

	return 0, nil
}

func (g *Vertex) compareTo(h *Vertex) int {
	return g.pzl.compareTo(h.pzl)
}
