package playingfield

import(
	"fmt"
)

func ExamplePlayingFieldImpl_String_emptyField() {
	field := New(5, 5)

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
	field := New(5,5,
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
	field := New(5,5,
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
	field := New(5,5,
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
	field := New(5,5,
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
	field := New(5,3,
		Position{2,2}, Position{0,0}, Position{4,1})

	fmt.Print(field)
	// Output:
	// *************
	// * - - O - - *
	// * - - - - O *
	// * O - - - - *
	// *************
}