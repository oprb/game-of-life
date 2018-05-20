package playingfield

import "fmt"

// PlayingField provides all necessary methods to play with an Game of Life playing field.
type PlayingField interface {
	// Update updates the state of the PlayingFIeld by applying the rules to all cells.
	Update()
	fmt.Stringer
	// CurrentState returns a view on the FieldState of the PlayingField. Changes to it may be reflected on the
	// FieldState of the PlayingField!
	CurrentState() *FieldState
	//FieldStateAt(time int) FieldState
}

// New provides a new PlayingField with  the cells  at the given Positions initially set to life.
func New(columns, rows int, initialPositions ...Position) PlayingField {
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
		adjacent.union(&neighbours)
	}

	positions.union(&adjacent)

	field.nextToUpdate = positions

	return &field
}

// FieldState represents the condition of all cells from a given PlayingField, where 'true' means the cell at the
// associated position is alive and 'false' indicates a dead cell.
type FieldState [][]bool