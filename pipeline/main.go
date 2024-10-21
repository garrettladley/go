package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	for num := range Take(done, Repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}

	fn := func() string { return "some fn call" }
	for num := range Take(done, RepeatFn(done, fn), 10) {
		fmt.Println(num)
	}

	start := time.Now()
	numFns := runtime.NumCPU()
	fmt.Printf("Spinning up %d fns.\n", numFns)
	fns := make([]<-chan string, numFns)
	fmt.Println("Fns:")

	longFn := func() string {
		time.Sleep(1 * time.Second)
		return fn() + " that takes a lloonngg time"
	}

	for i := 0; i < numFns; i++ {
		fns[i] = RepeatFn(done, longFn)
	}

	for result := range Take(done, FanIn(done, fns...), 10) {
		fmt.Printf("\t%s\n", result)
	}

	fmt.Printf("FanIn took %v.\n", time.Since(start))
}
