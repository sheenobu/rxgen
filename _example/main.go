package main

import (
	"fmt"

	"github.com/sheenobu/rxgen/rx"
)

func main() {

	b := rx.NewRxBool(true)

	sub := b.Subscribe()

	go func() {
		b.Set(false)
		b.Set(true)
		b.Set(false)

		sub.Close()
	}()

	for val := range sub.C {
		fmt.Printf("Got changed bool: %s\n", val)
	}
}
