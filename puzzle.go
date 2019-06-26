package slidingpuzzle

import "math/rand"

// Board ...
type Board [][]int

// Position ...
type Position []int

// Puzzle ...
type Puzzle struct {
	board        Board
	pos, prevPos Position
	rows, cols   int
}

// NewPuzzle ...
func NewPuzzle(m, n int) *Puzzle {
	pzl := &Puzzle{
		board:   make(Board, 0, m),
		pos:     Position{m - 1, n - 1},
		prevPos: Position{m - 1, n - 1},
		rows:    m,
		cols:    n,
	}

	var c int
	for i := 0; i < m; i++ {
		pzl.board = append(pzl.board, make([]int, 0, n))
		for j := 0; j < n; j++ {
			pzl.board[i] = append(pzl.board[i], c)
			c++
		}
	}

	return pzl
}

// solve ...
func (pzl *Puzzle) solve() {
	// TODO
}

// shuffle ...
func (pzl *Puzzle) shuffle(moves int) {
	var opts []Position
	for ; 0 < moves; moves-- {
		opts = pzl.options()
		pzl.slide(opts[rand.Intn(len(opts))])
	}
}

// slide ...
func (pzl *Puzzle) slide(opt Position) {
	copy(pzl.prevPos, pzl.pos)
	val := pzl.board[opt[0]][opt[1]]
	pzl.board[opt[0]][opt[1]] = pzl.board[pzl.pos[0]][pzl.pos[1]]
	pzl.board[pzl.pos[0]][pzl.pos[1]] = val
	copy(pzl.pos, opt)
}

// options ...
func (pzl *Puzzle) options() []Position {
	// Sorted options are up, left, right, down
	var opts []Position
	switch pzl.pos[0] {
	case 0:
		switch pzl.pos[1] {
		case 0:
			opts = []Position{pzl.right(), pzl.down()}
		case pzl.cols - 1:
			opts = []Position{pzl.left(), pzl.down()}
		default:
			opts = []Position{pzl.left(), pzl.right(), pzl.down()}
		}
	case pzl.rows - 1:
		switch pzl.pos[1] {
		case 0:
			opts = []Position{pzl.up(), pzl.right()}
		case pzl.cols - 1:
			opts = []Position{pzl.up(), pzl.left()}
		default:
			opts = []Position{pzl.up(), pzl.left(), pzl.right()}
		}
	default:
		switch pzl.pos[1] {
		case 0:
			opts = []Position{pzl.up(), pzl.right(), pzl.down()}
		case pzl.cols - 1:
			opts = []Position{pzl.up(), pzl.left(), pzl.down()}
		default:
			opts = []Position{pzl.up(), pzl.left(), pzl.right(), pzl.down()}
		}
	}

	n := len(opts)
	for i := 0; i < n; i++ {
		if opts[i] == nil {
			if i+1 < n {
				opts = append(opts[:i], opts[i+1:]...)
			}
			opts = opts[:i]
			n--
		}
	}

	// Don't allow previous move to be considered
	for i := range opts {
		if pzl.prevPos.compareTo(opts[i]) == 0 {
			if i+1 < n {
				return append(opts[:i], opts[i+1:]...)
			}
			return opts[:i]
		}
	}

	return opts
}

// up ...
func (pzl *Puzzle) up() Position {
	if pzl.pos[0] == 0 {
		return nil
	}
	return Position{pzl.pos[0] - 1, pzl.pos[1]}
}

// left ...
func (pzl *Puzzle) left() Position {
	if pzl.pos[1] == 0 {
		return nil
	}
	return Position{pzl.pos[0], pzl.pos[1] - 1}
}

// right ...
func (pzl *Puzzle) right() Position {
	if pzl.pos[1]+1 == pzl.cols {
		return nil
	}
	return Position{pzl.pos[0], pzl.pos[1] + 1}
}

// down ...
func (pzl *Puzzle) down() Position {
	if pzl.pos[0]+1 == pzl.rows {
		return nil
	}
	return Position{pzl.pos[0] + 1, pzl.pos[1]}
}

func (pos Position) compareTo(position Position) int {
	switch {
	case pos == nil:
		if position == nil {
			return 0
		}
		return 1
	case position == nil:
		return -1
	case pos[0] < position[0]:
		return -1
	case position[0] < pos[0]:
		return 1
	case pos[1] < position[1]:
		return -1
	case position[1] < pos[1]:
		return 1
	default:
		return 0
	}
}
