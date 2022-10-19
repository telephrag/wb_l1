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
		}(i) // i changes with each iteration so we need to copy it for calculation of each number
		// to avoid multiplying incorrect elements of nums
	}
	wg.Wait()  // to avoid exiting program preemptively
	close(res) // deadlock will occur if we range over opened channel

	for i := range res {
		fmt.Print(i, " ")
	}

	fmt.Println()
}
