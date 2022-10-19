package main

import "fmt"

func bSearch(s []int, v int) (index int) {
	i := len(s) / 2
	delta := i / 2
	for i >= 0 && i < len(s) {
		switch {
		case s[i] > v:
			i -= delta
		case s[i] < v:
			i += delta
		case s[i] == v:
			index = i
			return
		}
		delta = delta/2 + delta%2
	}

	return -1
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 31, 39, 47, 53, 59, 68, 73, 81, 90, 105, 107}

	fmt.Println(bSearch(input, 11))
	fmt.Println(bSearch(input, 1))
	fmt.Println(bSearch(input, 31))
	fmt.Println(bSearch(input, 105))
	fmt.Println(bSearch(input, 107))
	fmt.Println(bSearch(input, 0))
}
