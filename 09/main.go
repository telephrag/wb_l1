package main

import (
	"fmt"
	"os"
)

func writeNums(chNums chan int, input []int) {
	for num := range input {
		chNums <- num
	}
	// on closing of `chNums`, loop inside `writeSquares` will finish due to exhaustion of chNums
	close(chNums)
}

func writeSquares(chNums, chSqrs chan int) {
	for num := range chNums {
		chSqrs <- num * num
	}
	// on closing of `chSqrs`, loop inside `main` will finish due to exhaustion of `chSqrs`
	close(chSqrs)
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

	chNums := make(chan int, 2)
	chSqrs := make(chan int, 2)

	go writeNums(chNums, input)
	go writeSquares(chNums, chSqrs)

	for sqr := range chSqrs { // output squares of numbers until `chSqrs` is exhausted
		fmt.Fprint(os.Stdout, sqr, " ")
	}

	fmt.Println()
}
