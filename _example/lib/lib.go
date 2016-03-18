package lib

// generate some core types

//go:generate rxgen -name RxUint32 -type uint32 .
//go:generate rxgen -name RxInt32 -type int32 .

//go:generate rxgen -name RxUint64 -type uint64 .
//go:generate rxgen -name RxInt64 -type int64 .

//go:generate rxgen -name RxBool -type bool .

//go:generate rxgen -name RxString -type string .

// generate a custom type

//go:generate rxgen -name RxRect -type Rect .

// Rect represents a struct type that can be marked reactive.
type Rect struct {
	X int
	Y int
	W int
	H int
}
