package main

import (
	"fmt"
	"time"
)

// Despite placing several channels in our call to or that take various times to
// close, our channel that closes after one second causes the entire channel created by
// the call to Or to close.
//
// We achieve this terseness at the cost of additional goroutines—f(x)=⌊x/2⌋
func main() {
	start := time.Now()

	<-Or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}

func sig(after time.Duration) <-chan struct{} {
	c := make(chan struct{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}
