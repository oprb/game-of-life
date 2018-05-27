package playingfield

import "strings"

// PlayingFieldImpl provides a simple implementation of the PlayingField.
// It will be efficient, if sparsely occupied, but generates a bit overhead, if crowded by many living cells.
type PlayingFieldImpl struct {
	// total of columns
	columns int
	// total of rows
	rows int
	// number of iteration
	iteration int
	// current field state
	current FieldState
	previous []FieldState
	// positions that have to be regarded in the next iteration nd are likely to change their states
	nextToUpdate PositionSet
}

// CurrentState returns a view on the FieldState of the PlayingFieldImpl. Changes to it will be reflected on the
// FieldState of the PlayingFieldImpl!
func (field *PlayingFieldImpl) CurrentState() *FieldState {
	return &field.current
}

// Update updates the state of the PlayingFIeldImpl by applying the rules..
func (field *PlayingFieldImpl) Update() {
	var neighbours PositionSet
	// position that have to be inverted, i. e. cells that will live again or die
	invert := PositionSet{}
	positions := PositionSet{}

	for position := range field.nextToUpdate {
		neighbours = field.neighboursOf(position)
		living := field.cellsAlive(&neighbours)
		// cell is living but has too less or too much living neighbours to survive
		diesNext := field.current[position.Y][position.X] && (living < 2 || living > 3)
		// cell is dead but has exactly three living neighbours and will live next iteration
		willLiveNext := !field.current[position.Y][position.X] && living == 3

		if(diesNext || willLiveNext) {
			invert.add(position)
			positions.union(&neighbours)
		}
	}

	// invert the state all given positions
	for position := range invert {
		field.current[position.Y][position.X] = !field.current[position.Y][position.X]
	}

	field.nextToUpdate = positions
}

// cellsAlive provides the count of living cells among the given Positions.
func (field *PlayingFieldImpl) cellsAlive(positions *PositionSet) int {
	livingCells := 0
	for position := range *positions {
		if(field.current[position.Y][position.X]) {
			livingCells++
		}
	}

	return livingCells
}

// neighboursOf provides the neighbours for the cell at the given Position
func (field *PlayingFieldImpl) neighboursOf(position Position) PositionSet {
	// return the neighbours clockwise beginning at the top
	return PositionSet{
		field.remapPosition(Position{position.X, position.Y + 1}) : true,
		field.remapPosition(Position{position.X + 1, position.Y + 1}) : true,
		field.remapPosition(Position{position.X + 1, position.Y})  : true,
		field.remapPosition(Position{position.X + 1, position.Y - 1}) : true,
		field.remapPosition(Position{position.X, position.Y - 1}) : true,
		field.remapPosition(Position{position.X - 1, position.Y - 1}) : true,
		field.remapPosition(Position{position.X - 1, position.Y}) : true,
		field.remapPosition(Position{position.X - 1, position.Y + 1}) : true,
	}
}

// remapPosition remaps a given Position to the "associated" Position on the PlayingFieldImpl, it is called on.
// It serves as an helper function to ease the access to adjacent Positions of a Position at the borders of the
// PlayingFieldImpl.
func (field *PlayingFieldImpl) remapPosition(position Position) Position {
	return Position{ (position.X + field.columns) % field.columns, (position.Y + field.rows) % field.rows }
}

// String provides a string representation of the PlayingFieldImpl.
func (field *PlayingFieldImpl) String() string {
	const (
		living = "O "
		dead = "- "
	)

	borderPrinter := func(builder *strings.Builder) {
		borderElement := "*"
		for i := 0; i < field.columns * 2 + 3; i++ {
			builder.WriteString(borderElement)
		}
		builder.WriteString("\n")
	}

	builder := strings.Builder{}

	// print the upper boarder
	borderPrinter(&builder)
	// print all rows
	// (0,0) is in the bottom left, (n, m) is in the top right
	for rowIndex := field.rows -1; rowIndex >= 0; rowIndex-- {
		builder.WriteString("* ")

		for _, cellState := range field.current[rowIndex] {
			if cellState {
				builder.WriteString(living)
			} else {
				builder.WriteString(dead)
			}
		}
		builder.WriteString("*\n")
	}
	// print the lower boarder
	borderPrinter(&builder)

	return builder.String()
}
