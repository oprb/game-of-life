package playingfield

import (
	"testing"
	"errors"
)

func TestPlayingFieldImpl_CurrentState_withoutInitialPositionsSet(t *testing.T) {
	field, _ := newPlayingFieldImpl(5,5)

	given := field.CurrentState()
	expected := &FieldState{
		[]bool{false, false, false, false},
		[]bool{false, false, false, false},
		[]bool{false, false, false, false},
		[]bool{false, false, false, false},
		[]bool{false, false, false, false},}

	if given != expected {
		t.Fail()
	}
}

func TestPlayingFieldImpl_CurrentState_withInitialPositionsSet(t *testing.T) {
	field, _ := newPlayingFieldImpl(5,5, Position{2,4}, Position{1,1})

	given := field.CurrentState()
	expected := &FieldState{
		[]bool{false, false, true, false},
		[]bool{false, false, false, false},
		[]bool{false, false, false, false},
		[]bool{false, true, false, false},
		[]bool{false, false, false, false},}

	if given != expected {
		t.Fail()
	}
}

func TestPlayingFieldImpl_Update_withTooSmallInitialPositionsSetToCreateNewLife(t *testing.T) {
	field, _ := newPlayingFieldImpl(5,5, Position{2,4}, Position{1,1})
	field.Update()
	given := field.CurrentState()
	expected := &FieldState{
		[]bool{false, false, false, false},
		[]bool{false, false, false, false},
		[]bool{false, false, false, false},
		[]bool{false, false, false, false},
		[]bool{false, false, false, false},}

	if given != expected {
		t.Fail()
	}
}

func TestPlayingFieldImpl_Update_withInitialPositionsSet(t *testing.T) {
	field, _ := newPlayingFieldImpl(5,5, Position{0,0}, Position{0,1}, Position{1,1})

	given := field.CurrentState()
	// the expected initial setup
	expected := &FieldState{
		[]bool{false, false, false, false},
		[]bool{false, false, false, false},
		[]bool{false, false, false, false},
		[]bool{true, true, false, false},
		[]bool{true, false, false, false},}

	// 	test for the expected initial setup
	if given != expected {
		t.Fail()
	}

	field.Update()

	expected = &FieldState{
		[]bool{false, false, false, false},
		[]bool{false, false, false, false},
		[]bool{false, false, false, false},
		[]bool{false, false, false, false},
		[]bool{false, true, false, false},}
}

func newPlayingFieldImpl(columns, rows int, initialPositions ...Position) (*PlayingFieldImpl, error) {
	playingField, err := New(columns, rows, initialPositions...)

	if(err != nil) {
		return nil, err
	}

	 if playingFieldImpl, ok := playingField.(*PlayingFieldImpl); ok {
	 	return playingFieldImpl, nil
	 } else {
	 	return nil, errors.New("playingfield.New does no longer return a PlayingFieldImpl!")
	 }
}

