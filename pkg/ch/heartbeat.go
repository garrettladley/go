package ch

import (
	"time"
)

func HeartbeatOnInteval[T any](done <-chan struct{}, pulseInterval time.Duration, ch <-chan T) (<-chan struct{}, <-chan T) {
	heartbeat := make(chan struct{})
	results := make(chan T)
	go func() {
		defer close(heartbeat)
		defer close(results)
		pulse := time.Tick(pulseInterval)
		sendPulse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
		}
		sendResult := func(v T) {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case results <- v:
					return
				}
			}
		}
		for {
			select {
			case <-done:
				return
			case <-pulse:
				sendPulse()
			case v, ok := <-ch:
				if !ok {
					return
				}
				sendResult(v)
			}
		}
	}()

	return heartbeat, results
}

func HeartbeatOnStart[T any](done <-chan struct{}, ch <-chan T) (<-chan struct{}, <-chan T) {
	heartbeatStream := make(chan struct{}, 1)
	workStream := make(chan T)

	go func() {
		defer close(heartbeatStream)
		defer close(workStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-ch:
				if !ok {
					return
				}
				select {
				case heartbeatStream <- struct{}{}:
				default:
				}

				select {
				case workStream <- v:
				case <-done:
					return
				}
			}
		}
	}()
	return heartbeatStream, workStream
}
