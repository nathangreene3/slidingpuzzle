package slidingpuzzle

// Board ...
type Board [][]int

func (brd Board) validate() bool {
	var (
		m, n = brd.dimensions()
		mn   = m * n
	)
	switch {
	case m < 2, n < 2, mn < 6:
		return false
	}

	a := make([]bool, mn)
	for i := range brd {
		for _, val := range brd[i] {
			if a[val] {
				return false
			}
			a[val] = true
		}
	}

	return true
}

// entropy ...
func (brd Board) entropy() float64 {
	var (
		m, n = brd.dimensions()
		d    int
		e    float64
		k    int
	)
	if !brd.validate() {
		return 0
	}

	for i := range brd {
		for j := range brd[i] {
			d = brd[i][j] - k
			e += float64(d * d)
			k++
		}
	}

	return e / float64(m*n-1)
}

// dimensions returns the number of rows and columns of a board.
func (brd Board) dimensions() (int, int) {
	switch m := len(brd); m {
	case 0:
		return m, 0
	default:
		return m, len(brd[0])
	}
}

// copy ...
func (brd Board) copy() Board {
	m, n := brd.dimensions()
	b := make(Board, 0, m)
	for i := range brd {
		b = append(b, make([]int, n))
		copy(b[i], brd[i])
	}

	return b
}
