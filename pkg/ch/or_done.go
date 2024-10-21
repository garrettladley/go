package ch

// OrDone returns a channel that behaves like the provided channel, but also
// respects the done channel. The returned channel will close when either the
// provided channel or the done channel is closed.
//
// This function is useful to prevent goroutine leaks in case the done channel
// is closed before the provided channel.
//
// Also, it reduces the boilerplate of having a for select loop that checks
// if the context is done in every func.
func OrDone[T any](done <-chan struct{}, c <-chan T) <-chan T {
	relayStream := make(chan T)
	go func() {
		defer close(relayStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case relayStream <- v:
				case <-done:
					return
				}
			}
		}
	}()
	return relayStream
}
