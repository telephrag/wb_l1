package main

import "sync"

const loadPerThread = 24
const threadCount = 8

func processFlactuations(flsChan chan []float64, strg *ConcMap) {
	for fls := range flsChan {
		for _, f := range fls {
			fi := int64(f) // safe to do since it merely discards fraction part
			k := fi - fi%10
			strg.AppendAt(k, f) // appends slice at `k`, if doesn't exist creates a new one
		}
	}
}

func main() {
	// Divide input by chunks and pass them to threads.
	// Make each thread process numbers in it's chunk as described bellow:
	// Subsctract remainder from division of initial number by 10.
	// Store initial number into map using the result of subsctraction as a key.

	input := []float64{
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 33.6, 25.8, 19.3, 27.1, -4.2,
		-9.3, -15.7, -5.9, 1, -18.8, -21.5, -24.6, -19.71, -17.3, -28.51, -25, -24,
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 33.6, 25.8, 19.3, 27.1, -4.2,
		-9.3, -15.7, -5.9, 1, -18.8, -21.5, -24.6, -19.71, -17.3, -28.51, -25, -24,
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 33.6, 25.8, 19.3, 27.1, -4.2,
		-9.3, -15.7, -5.9, 1, -18.8, -21.5, -24.6, -19.71, -17.3, -28.51, -25, -24,
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 33.6, 25.8, 19.3, 27.1, -4.2,
		-9.3, -15.7, -5.9, 1, -18.8, -21.5, -24.6, -19.71, -17.3, -28.51, -25, -24,
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 33.6, 25.8, 19.3, 27.1, -4.2,
		-9.3, -15.7, -5.9, 1, -18.8, -21.5, -24.6, -19.71, -17.3, -28.51, -25, -24,
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 33.6, 25.8, 19.3, 27.1, -4.2,
		-9.3, -15.7, -5.9, 1, -18.8, -21.5, -24.6, -19.71, -17.3, -28.51, -25, -24,
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 33.6, 25.8, 19.3, 27.1, -4.2,
		-9.3, -15.7, -5.9, 1, -18.8, -21.5, -24.6, -19.71, -17.3, -28.51, -25, -24,
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 33.6, 25.8, 19.3, 27.1, -4.2,
		-9.3, -15.7, -5.9, 1, -18.8, -21.5, -24.6, -19.71, -17.3, -28.51, -25, -24,
	}
	l := len(input)

	workCh := make(chan []float64, 32)
	go func() {
		var left, right int = 0, loadPerThread % l
		for right < l {
			workCh <- input[left:right]
			left = right
			right = left + loadPerThread
		}
		if len(input[left:l]) != 0 {
			workCh <- input[left:l]
		}
		close(workCh) // close since there is nothing else to write
	}()

	var wg sync.WaitGroup
	wg.Add(threadCount)
	strg := (&ConcMap{}).Init()
	for i := 0; i < threadCount; i++ {
		go func() {
			processFlactuations(workCh, strg)
			wg.Done()
		}()
	}
	// work in `workCh` might not be processed before the end of the loop above
	// so we must explicitly wait before it is
	wg.Wait()

	strg.Print()
}
