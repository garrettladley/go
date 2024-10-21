package main

import (
	"fmt"

	"github.com/garrettladley/go/pkg/ch"
)

func main() {
	for v := range ch.Bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
}

func genVals() <-chan <-chan int {
	chanStream := make(chan (<-chan int))
	go func() {
		defer close(chanStream)
		for i := 0; i < 10; i++ {
			stream := make(chan int, 1)
			stream <- i
			close(stream)
			chanStream <- stream
		}
	}()
	return chanStream
}
