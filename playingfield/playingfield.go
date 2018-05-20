package playingfield

import "fmt"

type PlayingField interface {
	Update()
	fmt.Stringer
	//CurrentState() FieldState
	//FieldStateAt(time int) FieldState
	//Initialize(state FieldState)
}

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

type FieldState [][]bool