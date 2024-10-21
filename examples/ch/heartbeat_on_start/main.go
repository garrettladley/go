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

	heartbeat, results := ch.HeartbeatOnStart(done, intStream)

	for {
		select {
		case _, ok := <-heartbeat:
			if ok {
				fmt.Println("pulse")
			} else {
				return
			}
		case r, ok := <-results:
			if ok {
				fmt.Printf("results %v\n", r)
			} else {
				return
			}
		}
	}
}
