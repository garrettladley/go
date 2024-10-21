package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/garrettladley/go/pkg/ch"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	goodFner := func(i int) func() (int, error) { return func() (int, error) { return i, nil } }

	var retries int
	badFn := func() (int, error) {
		defer func() { retries++ }()
		if retries == 3 {
			return -1, nil
		}
		return 0, errors.New("error")
	}

	doWork, stream := ch.StewardWorkFn(done,
		[]func() (int, error){goodFner(0), goodFner(1), goodFner(2), badFn, goodFner(3)}...,
	)

	doWorkWithSteward := ch.NewSteward(1*time.Millisecond, doWork)

	doWorkWithSteward(done, 1*time.Hour)

	for intVal := range ch.Take(done, stream, 15) {
		fmt.Printf("Received: %v\n", intVal)
	}
}
