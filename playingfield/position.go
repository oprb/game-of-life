package playingfield

type Position struct {
	X int
	Y int
}

type PositionSet map[Position]bool

func (positions *PositionSet) add(position Position) bool {
	if _, alreadyIn := (*positions)[position]; alreadyIn {
		return true
	} else {
		(*positions)[position] = true
		return false
	}
}

func (positions *PositionSet) addAll(positionsToAdd ...Position) {
	for _, position := range positionsToAdd {
		positions.add(position)
	}
}

func (positions *PositionSet) union(newPositions *PositionSet) {
	for position := range *newPositions {
		positions.add(position)
	}
}
