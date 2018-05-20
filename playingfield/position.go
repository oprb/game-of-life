package playingfield

import (
	"fmt"
	"strings"
)

// Posution represents a position on the playing field with coordinates on the x and y axis.
type Position struct {
	X int
	Y int
}

// PositionSet represents a set of Positions.
type PositionSet map[Position]bool

// SortByPosition provides a way to sort Positions according to their coordinates.
type SortByPosition []Position

// add can be used to add a Position to a PositionSet.
func (positions *PositionSet) add(position Position) bool {
	if _, alreadyIn := (*positions)[position]; !alreadyIn {
		(*positions)[position] = true
		return true
	}

	return false
}

// addAll can be used to add  several Positions to a PositionSet.
func (positions *PositionSet) addAll(positionsToAdd ...Position) {
	for _, position := range positionsToAdd {
		positions.add(position)
	}
}

// union can be used to merge a PositionSet into the PositionSet, union is called on.
func (positions *PositionSet) union(newPositions *PositionSet) {
	for position := range *newPositions {
		positions.add(position)
	}
}

// String provides a string representation of a Position.
func (position Position) String() string {
	return fmt.Sprintf("[X: %v, Y: %v]", position.X, position.Y)
}

// Less is used to compare two positions by their coordinates while sorting.
func (positionsToSort *SortByPosition) Less(i, j int) bool {
	if ((*positionsToSort)[i].X < (*positionsToSort)[j].X ||
		(*positionsToSort)[i].X == (*positionsToSort)[j] .X && (*positionsToSort)[i].Y < (*positionsToSort)[j] .Y) {
		return true
	}

	return false
}

// asIterable provides all positions from a PositionSet as an iterable data structure.
func (positions *PositionSet) asIterable() []Position {
	iterable := make([]Position, len(*positions))

	for position := range *positions {
		iterable = append(iterable, position)
	}

	return iterable
}

// String provides a string representation of a PositionSet.
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