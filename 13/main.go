package main

import "fmt"

func swap(a, b *int) (*int, *int) { return b, a }

func main() {
	x, y := 1, 2
	a, b := &x, &y
	a, b = swap(a, b)
	fmt.Println(*a, *b)
}
