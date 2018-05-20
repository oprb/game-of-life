package playingfield

import "strings"

type PlayingFieldImpl struct {
	columns int
	rows int
	iteration int
	current FieldState
	previous []FieldState
	nextToUpdate PositionSet
}

func (field *PlayingFieldImpl) CurrentState() *FieldState {
	return &field.current
}

func (field *PlayingFieldImpl) Update() {
	var neighbours PositionSet
	// position that have to be inverted, i. e. cells that will live again or die
	invert := PositionSet{}
	updateNext := PositionSet{}

	for position := range field.nextToUpdate {
		neighbours = field.neighboursOf(position)
		living := field.living(&neighbours)
		// cell is living but has too less or too much living neighbours to survive
		diesNext := field.current[position.Y][position.X] && (living < 2 || living > 3)
		// cell is dead but has exactly three living neighbours and will live next iteration
		willLiveNext := !field.current[position.Y][position.X] && living == 3

		if(diesNext || willLiveNext) {
			invert.add(position)
			updateNext.union(&neighbours)
		}
	}

	for position := range invert {
		field.current[position.Y][position.X] = !field.current[position.Y][position.X]
	}

	field.nextToUpdate = updateNext
}

func (field *PlayingFieldImpl) invert(position Position) {
	field.current[position.Y][position.X] = !field.current[position.Y][position.X]
}

func (field *PlayingFieldImpl) living(positions *PositionSet) int {
	livingCells := 0
	for position := range *positions {
		if(field.current[position.Y][position.X]) {
			livingCells++
		}
	}

	return livingCells
}

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

// utility function to ease the access of neighbours from positions at the borders of the field
func (field *PlayingFieldImpl) remapPosition(position Position) Position {
	return Position{ (position.X + field.columns) % field.columns, (position.Y + field.rows) % field.rows }
}

func (field *PlayingFieldImpl) String() string {
	borderPrinter := func(builder *strings.Builder) {
		borderElement := "*"
		for i := 0; i < field.columns * 2 + 3; i++ {
			builder.WriteString(borderElement)
		}
		builder.WriteString("\n")
	}

	living := "O "
	dead := "- "
	builder := strings.Builder{}

	// print the upper boarder
	borderPrinter(&builder)
	// prinnt all rows
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
