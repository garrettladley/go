package main

import (
	"fmt"
	"time"

	"github.com/garrettladley/go/pkg/ch"
	"golang.org/x/exp/rand"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	fn := func() string {
		duration := time.Duration(rand.Intn(5)+1) * time.Second
		time.Sleep(duration)
		return "hello"
	}

	start := time.Now()
	result := ch.Replicated(done, 10, fn)
	elapsed := time.Since(start)

	fmt.Printf("result: %v\n", result)
	fmt.Printf("elapsed: %v\n", elapsed)
}
