package ch

func Generator[T any](done <-chan struct{}, vals ...T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, v := range vals {
			select {
			case <-done:
				return
			case ch <- v:
			}
		}
	}()
	return ch
}
