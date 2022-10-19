package main

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

func workProvider(ctx context.Context, work chan []byte) {
	b := md5.Sum([]byte(fmt.Sprint(rand.Int63())))
	hash := b[:]

	for {
		select {
		case <-ctx.Done():
			close(work) // workers are ranging over `work` so,
			return      //we need to close it to prevent deadlock
		default:
			b := md5.Sum(hash)
			hash = b[:]
			work <- b[:]
			// comment above and uncomment bellow to see that program runs correctly
			// work <- []byte("a very long debug sentence 1234567890")
			time.Sleep(time.Nanosecond * time.Duration(rand.Int63()%100))
		}
	}
}

var workersRunning int32 // can be made local and passed to each worker but whatever

func workConsumer(id int, work chan []byte) {
	atomic.AddInt32(&workersRunning, 1)        // tracking amount of active workers
	defer atomic.AddInt32(&workersRunning, -1) // is needed for gracefully shutting down program

	for b := range work { // worker will be ranging over `work` while it's open
		doubleHash := md5.Sum(b)
		fmt.Fprintf(os.Stdout, "%d: %s\n", id, base64.StdEncoding.EncodeToString(doubleHash[:]))
	}
}

func main() {
	var workerCount int
	flag.IntVar(&workerCount, "wc", 1, "Amount of workers")
	flag.Parse()

	// context is used to signal `workProvider` to shutdown
	ctx, cancel := context.WithCancel(context.Background())
	work := make(chan []byte, 32)

	go workProvider(ctx, work)

	for i := 0; i < workerCount; i++ {
		go workConsumer(i, work)
	}

	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, syscall.SIGINT, syscall.SIGTERM)
	<-interupt

	cancel() // stop providing work on ^C

	// In more complex program, a work provided by `workProvider` if left incomplete
	// may leave the whole program in a broken state.
	// In this case we need to rollback what was already complete or get work we've already
	// received done. The later solution is easier to write and maintain.
	// This spreads on client side code since an error when workers shutdown before
	// requested work could be completed is not present of course,
	// provided that error can not occur during performance of the received work.

	// As a result of our choice we wait until `work` is exhausted
	// and all the remaining work has been done.
	for atomic.LoadInt32(&workersRunning) != 0 {
	}
}
