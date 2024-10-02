package main

import "context"

// orDone returns a channel that behaves like the provided channel, but also
// respects the context. If the context is canceled, the returned channel will
// be closed.
//
// This function is useful to prevent goroutine leaks in case the context is
// canceled before the goroutine is done.
//
// Also, it reduces the boilerplate of having a for select loop that checks
// if the context is done in every func.
func orDone[T any](ctx context.Context, c <-chan T) <-chan T {
	relayStream := make(chan T)
	go func() {
		defer close(relayStream)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case relayStream <- v:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return relayStream
}
