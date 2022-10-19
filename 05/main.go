package main

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	b := md5.Sum([]byte(fmt.Sprint(rand.Int63())))
	hash := b[:]

	ch := make(chan string, 16)

	go func() {
		for {
			b := md5.Sum(hash) // creating some data to send via channel
			hash = b[:]        // comment this line to see that message arrive uncorrupted
			ch <- base64.StdEncoding.EncodeToString(b[:])
		}
	}()

	go func() {
		for msg := range ch {
			fmt.Fprintf(os.Stdout, "Chan read: %s\n", msg)
		}
	}()

	// Tieing program shutdown to automatic context competion on timeout.
	// Same can be done using `time.After()` which is used internally
	// inside `context.WithTimeout`
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
