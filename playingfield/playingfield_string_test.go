package playingfield

import(
	"fmt"
)

func ExamplePlayingFieldImpl_String_emptyField() {
	field, _ := New(5, 5)

	fmt.Print(field)
	// Output:
	// *************
	// * - - - - - *
	// * - - - - - *
	// * - - - - - *
	// * - - - - - *
	// * - - - - - *
	// *************
}

func ExamplePlayingFieldImpl_String_positionAtZeroZeroSet() {
	field, _ := New(5,5,
		Position{0,0})

	fmt.Print(field)
	// Output:
	// *************
	// * - - - - - *
	// * - - - - - *
	// * - - - - - *
	// * - - - - - *
	// * O - - - - *
	// *************
}

func ExamplePlayingFieldImpl_String_positionAtFourFourSet() {
	field, _ := New(5,5,
		Position{4,4})

	fmt.Print(field)
	// Output:
	// *************
	// * - - - - O *
	// * - - - - - *
	// * - - - - - *
	// * - - - - - *
	// * - - - - - *
	// *************
}

func ExamplePlayingFieldImpl_String_positionAtFourZeroSet() {
	field, _ := New(5,5,
		Position{4,0})

	fmt.Print(field)
	// Output:
	// *************
	// * - - - - - *
	// * - - - - - *
	// * - - - - - *
	// * - - - - - *
	// * - - - - O *
	// *************
}

func ExamplePlayingFieldImpl_String_positionAtTwoThreeSet() {
	field, _ := New(5,5,
		Position{2,3})

	fmt.Print(field)
	// Output:
	// *************
	// * - - - - - *
	// * - - O - - *
	// * - - - - - *
	// * - - - - - *
	// * - - - - - *
	// *************
}

func ExamplePlayingFieldImpl_String_multiplePositionsSet() {
	field, _ := New(5,3,
		Position{2,2}, Position{0,0}, Position{4,1})

	fmt.Print(field)
	// Output:
	// *************
	// * - - O - - *
	// * - - - - O *
	// * O - - - - *
	// *************
}