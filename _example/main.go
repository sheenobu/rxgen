package main

import (
	"fmt"
	"time"

	"github.com/sheenobu/rxgen/rx"
)

func main() {

	b := rx.NewBool(true)

	sub := b.Subscribe()

	go func() {
		b.Set(false)
		<-time.After(100 * time.Millisecond)
		b.Set(true)
		<-time.After(100 * time.Millisecond)
		b.Set(false)
		<-time.After(100 * time.Millisecond)

		sub.Close()
	}()

	for val := range sub.C {
		fmt.Printf("Got changed bool: %v\n", val)
	}
}
