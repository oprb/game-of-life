package playingfield

import (
	"fmt"
	"errors"
	"strings"
)

// PlayingField provides all necessary methods to play with an Game of Life playing field.
type PlayingField interface {
	// Update updates the state of the PlayingFIeld by applying the rules to all cells.
	Update()
	fmt.Stringer
	// CurrentState returns a view on the FieldState of the PlayingField. Changes to it may be reflected on the
	// FieldState of the PlayingField!
	CurrentState() FieldState
	//FieldStateAt(time int) FieldState
}

// New provides a new PlayingField with  the cells  at the given Positions initially set to life.
// Duplicated positions will be ignored.
func New(columns, rows int, initialPositions ...Position) (PlayingField, error) {
	if err := checkForInvalidPositions(columns, rows, initialPositions...); err != nil {
		return nil, err
	}

	// initialize a playing field
	field := PlayingFieldImpl{
		columns: columns,
		rows: rows,
		current: make([][]bool, rows),
	}

	// add rows, each with length 'columns'
	for i := range field.current {
		field.current[i] = make([]bool, columns)
	}

	// set the cells at the given positions, they + their neighbours have to be updated in the next iteration
	positions := PositionSet{}
	adjacent := PositionSet{}
	for _, position := range initialPositions {
		field.current[position.Y][position.X] = true

		positions.add(position)
		neighbours := field.neighboursOf(position)
		adjacent.union(neighbours)
	}

	positions.union(adjacent)

	field.nextToUpdate = positions

	return &field, nil
}

// checkForInvalidPositions checks for coordinates among the given Positions, that does not fit onto a PlayingField
// of size 'columns' X 'rows' and returns an error listing the invalid positions, in case such were found.
func checkForInvalidPositions(columns, rows int, positions ...Position) error {
	nonValidPositions := []Position{}
	// collect the invalid positions
	for _, position := range positions {
		if position.X < 0 || position.Y < 0 || position.X >= columns || position.Y >= rows {
			nonValidPositions = append(nonValidPositions, position)
		}
	}

	// use a string builder to create the listing of invalid positions and return an error
	if length := len(nonValidPositions); length != 0 {
		builder := strings.Builder{}
		builder.WriteString("playingField.New: Position/s with invalid coordinate/s were given: ")
		if length > 1 {
			for i := 0; i < length - 1; i++ {
				builder.WriteString(fmt.Sprintf("%v, ", nonValidPositions[i]))
			}
		}
		builder.WriteString(fmt.Sprintf("%v", nonValidPositions[length - 1]))

		return errors.New(builder.String())
	}

	return nil
}

// FieldState represents the condition of all cells from a given PlayingField, where 'true' means the cell at the
// associated position is alive and 'false' indicates a dead cell.
type FieldState [][]bool