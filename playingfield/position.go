package playingfield

import (
	"fmt"
	"strings"
)

type Position struct {
	X int
	Y int
}

type PositionSet map[Position]bool

type SortByPosition []Position

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

func (position Position) String() string {
	return fmt.Sprintf("[X: %v, Y: %v]", position.X, position.Y)
}


func (positions *SortByPosition) Len() int {
	return len([]Position(*positions))
}

func (positions *SortByPosition) Swap(i, j int) {

}

func (positionsToSort *SortByPosition) Less(i, j int) bool {
	if ((*positionsToSort)[i].X < (*positionsToSort)[j].X ||
		(*positionsToSort)[i].X == (*positionsToSort)[j] .X &&  (*positionsToSort)[i].Y < (*positionsToSort)[j] .Y) {
		return true
	}

	return false
}

func (positions *PositionSet) asIterable() []Position {
	iterable := make([]Position, len(*positions))

	for position := range *positions {
		iterable = append(iterable, Position{position.X, position.Y})
	}

	return iterable
}

func (positions *PositionSet) String() string {
	builder := strings.Builder{}
	//positionsToSort := positions.asIterable()
	//sort.Sort(SortByPosition(positionsToSort)
	builder.WriteString("{\n")
	for position := range *positions {
		builder.WriteString(fmt.Sprintf(" %v,\n", position))
	}
	builder.WriteString("}")

	return builder.String()
}