package ch

func Repeat[T any](done <-chan struct{}, values ...T) <-chan T {
	valueStream := make(chan T)
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

func RepeatFn[T any](done <-chan struct{}, fns ...func() T) <-chan T {
	valueStream := make(chan T)
	go func() {
		defer close(valueStream)
		for {
			for _, fn := range fns {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}
	}()
	return valueStream
}
