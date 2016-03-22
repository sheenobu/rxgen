package lib

// generate a custom type

//go:generate rxgen -name RxRect -type Rect .

// Rect represents a struct type that can be marked reactive.
type Rect struct {
	X int
	Y int
	W int
	H int
}
