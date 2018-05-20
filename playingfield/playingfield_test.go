package playingfield

import "testing"

func TestPlayingFieldImpl_CurrentState_withoutInitialPositionsSet(t *testing.T) {
	field := New(5,5).(*PlayingFieldImpl)

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
	field := New(5,5, Position{2,4}, Position{1,1}).(*PlayingFieldImpl)

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
	field := New(5,5, Position{2,4}, Position{1,1}).(*PlayingFieldImpl)
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
	field := New(5,5, Position{0,0}, Position{0,1}, Position{1,1}).(*PlayingFieldImpl)

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



