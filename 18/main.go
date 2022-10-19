package main

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type Counter struct {
	c int32
}

func (c *Counter) Inc() {
	atomic.AddInt32(&c.c, 1)
}

func (c *Counter) Get() int32 {
	return atomic.LoadInt32(&c.c)
}

const workerCount = 8

func main() {

	c := Counter{}

	for i := 0; i < workerCount; i++ {
		go func(workerID int) {
			r := rand.New(rand.NewSource(time.Now().UnixMicro()))
			b := md5.Sum([]byte(fmt.Sprint(r.Int63())))
			hash := b[:]
			for {
				b = md5.Sum(hash)
				hash = b[:]
				c.Inc()
				fmt.Printf("%d %-15s %d\n", workerID, base64.StdEncoding.EncodeToString(hash), c.Get())
			}
		}(i)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond)
	for range ctx.Done() {
	}
}
