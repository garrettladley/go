package main

import "sync"

func FanIn[T any](done <-chan struct{}, channels ...<-chan T) <-chan T {
	var (
		wg                sync.WaitGroup
		multiplexedStream = make(chan T)
	)

	multiplex := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	// select from all the channels
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// wait for all the reads to complete
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
