package main

import (
	"fmt"
	"time"

	"github.com/garrettladley/go/pkg/ch"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 0; i < 10; i++ {
			intStream <- i
			time.Sleep(1 * time.Second)
		}
	}()

	const timeout = 2 * time.Second
	heartbeat, results := ch.HeartbeatOnInteval(done, timeout/2, intStream)

	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if !ok {
				return
			}
			fmt.Printf("results %v\n", r)
		case <-time.After(timeout):
			fmt.Println("worker goroutine is not healthy!")
			return
		}
	}
}
