package main

func Take[T any](done <-chan struct{}, valueStream <-chan T, num uint) <-chan T {
	takeStream := make(chan T)
	go func() {
		defer close(takeStream)
		n := int(num)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}
