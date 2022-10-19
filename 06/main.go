package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

var runningCount int32

// All of the ways to stop goroutine using context.
// +3
func ctx() {
	parentCtx := context.Background()

	// Comment two out of three bellow.

	ctx, cc := context.WithCancel(parentCtx)
	cc()

	// ctx, _ := context.WithTimeout(parentCtx, 0)

	// ctx, _ := context.WithDeadline(parentCtx, time.Now())

	go func() {
		// Go doesn't provide a way to track amount of goroutines doing some job
		// and not simply sleeping in pool with a standard library so,
		// we'll have to do it ourselves with these two lines bellow.
		atomic.AddInt32(&runningCount, 1)
		defer atomic.AddInt32(&runningCount, -1)

		for range ctx.Done() {
		}
	}()
}

// Stop goroutine by returning from it.
// +1
func ret() {
	go func() {
		atomic.AddInt32(&runningCount, 1)
		defer atomic.AddInt32(&runningCount, -1)
	}()
}

// Stop goroutine via closing channel it's reading from.
// Idk if I should include it since, we're doing pretty much the same as in `stopContext()`
// +2
func chanRange() {

	ch := make(chan struct{}, 1)

	go func() {
		atomic.AddInt32(&runningCount, 1)
		defer atomic.AddInt32(&runningCount, -1)

		for range ch {
		}
	}()

	close(ch)
}

// Stop goroutine by panicing inside of it.
// +1
func panicInside() {
	go func() {
		atomic.AddInt32(&runningCount, 1)
		defer atomic.AddInt32(&runningCount, -1)

		defer func() {
			recover() // gotta recover to show you other ways
		}()
		panic("oops...")
	}()
}

// Stop goroutine using built in `time.After()`
// Again, not sure if this should be included since we are ranging over channel
// that will be closed later.
// +1
func timeout() {
	ch := time.After(0)

	go func() {
		atomic.AddInt32(&runningCount, 1)
		defer atomic.AddInt32(&runningCount, -1)

		for range ch {
		}
	}()
}

// Stop via reading from channel inside `select`.
// Reads can occur periodically or continiously.
// +1
func slct() {
	ch := make(chan struct{}, 1)

	go func() {
		atomic.AddInt32(&runningCount, 1)
		defer atomic.AddInt32(&runningCount, -1)

		for {
			select {
			case <-ch:
				return
			default:
				time.Sleep(time.Millisecond * 50) // can be commented out
			}
		}
	}()

	time.Sleep(time.Millisecond * 75)
	ch <- struct{}{}
}

func main() {
	// (ctx cancel) x (cancel call, timeout, deadline) V
	// return V
	// closing of channel being ranged over V
	// panic inside goroutine itself V
	// time.After() V
	// periodic polling V
	// stop of the main goroutine
	// (fatal) x (in goroutine itself, somewhere else)

	ctx()
	ret()
	chanRange()
	panicInside()
	timeout()
	slct()

	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, syscall.SIGINT, syscall.SIGTERM)
	<-interupt

	// Stop goroutine via stopping `main` goroutine
	// +1
	go func() {
		for {
		}
	}()

	// Stop goroutine via calling fatal inside of it.
	// +1
	// go func() {
	// 	log.Fatal()
	// }()

	// 11 ways total to stop goroutine.

	time.Sleep(time.Millisecond * 100) // let all of the goroutines to finish
	fmt.Printf("\n goroutines running: %d\n", runningCount)
}
