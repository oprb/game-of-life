package playingfield

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
		case field.current[position.X][position.Y] && (living < 2 || living > 3):
			invert.add(position)
			updateNext.union(&neighbours)
			// cell is dead but has exactly three living neighbours and will live again
		case !field.current[position.X][position.Y] && living == 3:
			invert.add(position)
			updateNext.union(&neighbours)
		}
	}

	for position := range invert {
		field.current[position.X][position.Y] = !field.current[position.X][position.Y]
	}

	field.nextToUpdate = updateNext
}

func (field *PlayingFieldImpl) Initialize(columns, rows int, positions PositionSet) {
	field.current = make([][]bool, columns)

	for i := range field.current {
		field.current[i] = make([]bool, rows)
	}

	for position := range positions {
		field.current[position.X][position.Y] = true
	}

	field.nextToUpdate = positions
}

func (field *PlayingFieldImpl) invert(position Position) {
	field.current[position.X][position.Y] = !field.current[position.X][position.Y]
}

func (field *PlayingFieldImpl) living(positions *PositionSet) int {
	livingCells := 0
	for position := range *positions {
		if(field.current[position.X][position.Y]) {
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

func (field *PlayingFieldImpl) remapPosition(position Position) Position {
	return Position{ (position.X + field.columns) % field.columns, (position.Y + field.rows) % field.rows }
}

func New(columns, rows int, initialPositions ...Position) PlayingField {
	// initialize a playing field
	field := PlayingFieldImpl{
		columns: columns,
		rows: rows,
		current: make([][]bool, columns),
	}

	// add rows
	for i := range field.current {
		field.current[i] = make([]bool, rows)
	}

	// set the cells at the given positions, them + their neighbours have to be updated in the next iteration
	positions := PositionSet{}
	adjacent := PositionSet{}
	for _, position := range initialPositions {
		field.current[position.X][position.Y] = true

		positions.add(position)
		neighbours := field.neighboursOf(position)
		adjacent.union(&neighbours)
	}

	positions.union(&adjacent)

	field.nextToUpdate = positions

	return &field
}

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

type FieldState [][]bool