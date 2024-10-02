package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	gophers := make(chan string)

	go produce(ctx, &wg, gophers)

	wg.Add(1)
	go consume(ctx, &wg, gophers)

	wg.Wait()
}

func consume(ctx context.Context, wg *sync.WaitGroup, gophers <-chan string) {
	defer wg.Done()
	for val := range orDone(ctx, gophers) {
		fmt.Println(val)
	}
}

func produce(ctx context.Context, wg *sync.WaitGroup, gophers chan<- string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case gophers <- "go":
		}
	}
}
