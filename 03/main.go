package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	res := make(chan int, len(nums))

	var wg sync.WaitGroup
	wg.Add(len(nums))
	for i := range nums {
		go func(index int) {
			res <- nums[index] * nums[index]
			wg.Done()
		}(i) // `i`` changes with each iteration so we need to copy it for calculation of each number
		// to avoid multiplying incorrect elements of `nums``
	}
	wg.Wait()  // to avoid closing `res` preemptively
	close(res) // deadlock will occur if we range over opened channel

	done := make(chan struct{}, 1) // `struct{}` as chan type for less memory allocations
	defer close(done)

	var sum int
	go func() {
		for i := range res { // will be ranging until `res` is closed
			sum += i
		}
		done <- struct{}{} // signalling that we finished the work
	}()

	<-done // dellaying output till summation is done

	fmt.Println(sum)
}
