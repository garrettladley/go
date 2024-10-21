package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/garrettladley/go/pkg/ch"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	for num := range ch.Take(done, ch.Repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}

	fmt.Print("\n")

	for num := range ch.Take(done, ch.RepeatFn(done, func() string { return "0" }, func() string { return "1" }), 10) {
		fmt.Printf("%v ", num)
	}

	start := time.Now()
	numFns := runtime.NumCPU()
	fmt.Printf("Spinning up %d fns.\n", numFns)
	fns := make([]<-chan string, numFns)
	fmt.Println("Fns:")

	longFn := func() string {
		time.Sleep(1 * time.Second)
		return "fn that takes a lloonngg time"
	}

	for i := 0; i < numFns; i++ {
		fns[i] = ch.RepeatFn(done, longFn)
	}

	for result := range ch.Take(done, ch.FanIn(done, fns...), 10) {
		fmt.Printf("\t%s\n", result)
	}

	fmt.Printf("FanIn took %v.\n", time.Since(start))
}
