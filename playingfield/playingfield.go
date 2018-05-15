package playingfield

import(
	"strings"
)

type PlayingField interface {
	Update()
	//CurrentState() FieldState
	//FieldStateAt(time int) FieldState
	//Initialize(state FieldState)
}

type PlayingFieldImpl struct {
	columns int
	rows int
	iteration int
	current FieldState
	previous []FieldState
	nextToUpdate PositionSet
}


func (field *PlayingFieldImpl) Update() {
	var neighbours PositionSet
	// position that have to be inverted, i. e. cells that will live again or die
	invert := PositionSet{}
	updateNext := PositionSet{}

	for position := range field.nextToUpdate {
		neighbours = field.neighboursOf(position)
		field.invert(position)

		living := field.living(&neighbours)

		switch {
		// cell is living but has too less or too much living neighbours to survive
		case field.current[position.Y][position.X] && (living < 2 || living > 3):
			invert.add(position)
			updateNext.union(&neighbours)
			// cell is dead but has exactly three living neighbours and will live again
		case !field.current[position.Y][position.X] && living == 3:
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
		borderElement := "**"
		for i := 0; i < field.columns + 1; i++ {
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
		builder.WriteString("*")

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

	// set the cells at the given positions, them + their neighbours have to be updated in the next iteration
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