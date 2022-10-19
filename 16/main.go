package main

import "fmt"

func partition(arr []int) (left []int, right []int) {
	// pivo := (0 + len(arr) - 1) / 2
	pivo := len(arr) - 1
	// fmt.Printf("pivo: %d\n", pivo)

	i := -1
	for j := 0; j < pivo; j++ {
		if arr[j] < arr[pivo] {
			i++
			arr[j], arr[i] = arr[i], arr[j]
			// fmt.Printf("j: %2d; i: %2d; arr: %v\n", j, i, arr) // debug
		}
	}

	arr[i+1], arr[pivo] = arr[pivo], arr[i+1]

	// fmt.Printf("end: %v; %d; %v\n\n", arr[:i+1], arr[i+1], arr[i+2:]) // debug

	return arr[:i+1], arr[i+2:]
}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := partition(arr)
	quicksort(left)
	quicksort(right)

	return arr
}

func main() {
	input := []int{7, 3, 11, 22, 17, 18, 19, 90, 37, 63, 100, 8}
	fmt.Println(quicksort(input))

	input = []int{5, 78, 516, 29, 0, 5, 72, 17, 520, 3}
	fmt.Println(quicksort(input))

}
