package ch

func Bridge[T any](done <-chan struct{}, chanStream <-chan <-chan T) <-chan T {
	valStream := make(chan T)
	go func() {
		defer close(valStream)
		for {
			var stream <-chan T
			select {
			case maybeStream, ok := <-chanStream:
				if ok == false {
					return
				}
				stream = maybeStream
			case <-done:
				return
			}
			for val := range OrDone(done, stream) {
				select {
				case valStream <- val:
				case <-done:
				}
			}
		}
	}()
	return valStream
}
