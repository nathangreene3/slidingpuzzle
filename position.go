package slidingpuzzle

import "sort"

// Position ...
type Position []int

// compareTo ...
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

func (pos Position) copy() Position {
	p := make(Position, len(pos))
	copy(p, pos)
	return p
}

func sortPositions(pos []Position) {
	sort.Slice(pos, func(i, j int) bool { return pos[i].compareTo(pos[j]) < 0 })
}

func stableSortPositions(pos []Position) {
	sort.SliceStable(pos, func(i, j int) bool { return pos[i].compareTo(pos[j]) < 0 })
}
