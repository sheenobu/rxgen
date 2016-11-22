package lib

// generate a custom type

// Rect represents a struct type that can be marked reactive.
// +gen rx
type Rect struct {
	X int
	Y int
	W int
	H int
}
