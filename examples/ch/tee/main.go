package main

import (
	"fmt"

	"github.com/garrettladley/go/pkg/ch"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	out1, out2 := ch.Tee(done, ch.Take(done, ch.Repeat(done, 1, 2), 4))
	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
}
