package ch

import (
	"sync"
)

func Replicated[R any](done <-chan struct{}, n uint, fn func() R) R {
	var (
		replicatedDone = make(chan struct{})
		xDone          = Or(done, replicatedDone)

		wg     sync.WaitGroup
		result = make(chan R, 1)
	)

	for i := uint(0); i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-xDone:
				return
			case result <- fn():
			}
		}()
	}

	go func() {
		wg.Wait()
		close(replicatedDone)
		close(result)
	}()

	select {
	case <-done:
		var zeroR R
		return zeroR
	case res := <-result:
		return res
	}
}
