package slidingpuzzle

// Node ...
type Node struct {
	pzl   *Puzzle
	nbors []*Node
}

// NewGraph ...
func NewGraph(pzl *Puzzle) *Node {
	return &Node{pzl: pzl, nbors: make([]*Node, 0, 256)}
}
