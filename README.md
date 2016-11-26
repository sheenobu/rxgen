# rxgen

rxgen is a gen typewriter for generating reactive code

## installation

    gen add github.com/sheenobu/rxgen

## Simple Usage

Outputs type `RxRect`:

```go
// +gen rx
type Rect struct {
	X int
	Y int
	W int
	H int
}
```

## Builtin types

you can annotate builtin types with a type alias and an underscore:

```go
// +gen rx:"Builtin[uint]"
type Uint_ uint
```

`Uint_` will be ignored and the output name will be `Uint` and wrap `uint`. Additionally,
you can use build tags to exclude `Uint_` from being included.


## Usage

```go

import (
   "lib"
)

r := NewUint(1)
r2 := NewRxRect(Rect{X: 0, Y: 0, W: 0, H: 0})

// Get the value
v := r.Get() // 1

// Create a new subscription
c := r.Subscribe()

go func() {
	r.Set(1234)  // blocks, triggers loop below
	r.Set(4321)  // blocks, triggers loop below

	c.Close() // closes subscription, quits loop below
}()

for val := range c.C {
	fmt.Printf("Integer changed value: %d\n", val)
}
```
