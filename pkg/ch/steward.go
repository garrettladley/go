package ch

import (
	"log"
	"time"
)

type StartGoroutineFn func(
	done <-chan struct{},
	pulseInterval time.Duration,
) (heartbeat <-chan struct{})

func NewSteward(
	timeout time.Duration,
	startGoroutine StartGoroutineFn,
) StartGoroutineFn {
	return func(
		done <-chan struct{},
		pulseInterval time.Duration,
	) <-chan struct{} {
		heartbeat := make(chan struct{})
		go func() {
			defer close(heartbeat)
			var wardDone chan struct{}
			var wardHeartbeat <-chan struct{}
			startWard := func() {
				wardDone = make(chan struct{})
				wardHeartbeat = startGoroutine(Or(wardDone, done), timeout/2)
			}
			startWard()
			pulse := time.Tick(pulseInterval)
		monitorLoop:
			for {
				timeoutSignal := time.After(timeout)
				for {
					select {
					case <-pulse:
						select {
						case heartbeat <- struct{}{}:
						default:
						}
					case <-wardHeartbeat:
						continue monitorLoop
					case <-timeoutSignal:
						log.Println("steward: ward unhealthy; restarting")
						close(wardDone)
						startWard()
						continue monitorLoop
					case <-done:
						return
					}
				}
			}
		}()
		return heartbeat
	}
}

func StewardWorkFn[T any](done <-chan struct{}, fns ...func() (T, error)) (StartGoroutineFn, <-chan T) {
	valChanStream := make(chan (<-chan T))
	valStream := Bridge(done, valChanStream)
	doWork := func(
		done <-chan struct{},
		pulseInterval time.Duration,
	) <-chan struct{} {
		valStream := make(chan T)
		heartbeat := make(chan struct{})
		go func() {
			defer close(valStream)
			select {
			case valChanStream <- valStream:
			case <-done:
				return
			}

			pulse := time.Tick(pulseInterval)

			for {
			valueLoop:
				for _, fn := range fns {
					val, err := fn()
					if err != nil {
						log.Printf("steward: %v\n", err)
						return
					}

					for {
						select {
						case <-pulse:
							select {
							case heartbeat <- struct{}{}:
							default:
							}
						case valStream <- val:
							continue valueLoop
						case <-done:
							return
						}
					}
				}
			}
		}()
		return heartbeat
	}
	return doWork, valStream
}
