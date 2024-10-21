package main

import (
	"fmt"

	"github.com/garrettladley/go/pkg/ch"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	for v := range ch.Generator(done, 1, 2, 3, 4, 5) {
		fmt.Printf("%d ", v)
	}
}
