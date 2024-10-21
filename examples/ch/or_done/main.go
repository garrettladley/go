package main

import (
	"fmt"
	"sync"

	"github.com/garrettladley/go/pkg/ch"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	var wg sync.WaitGroup

	gophers := make(chan string)

	go produce(done, &wg, gophers)

	wg.Add(1)
	go consume(done, &wg, gophers)

	wg.Wait()
}

func consume(done <-chan struct{}, wg *sync.WaitGroup, gophers <-chan string) {
	defer wg.Done()
	for val := range ch.OrDone(done, gophers) {
		fmt.Println(val)
	}
}

func produce(done <-chan struct{}, wg *sync.WaitGroup, gophers chan<- string) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		case gophers <- "go":
		}
	}
}
