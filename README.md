# rxgen

golang library for generating reactive code

## installation

    go install github.com/sheenobu/rxgen

## example

lib/lib.go:

    package lib

    //go:generate rxgen -type uint -name RxUint .

main.go:

    import (
       "lib"
    )

    r := lib.NewRxUint(1)

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

