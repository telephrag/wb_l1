package main

import "fmt"

func intersect(a, b []int) []int {
	m := make(map[int]int)

	for _, v := range a {
		m[v]++
	}

	for _, v := range b {
		m[v]++
	}

	fmt.Println(m)

	res := make([]int, 0)
	i := 0
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
		i++
	}

	return res
}

func main() {
	a := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	b := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}

	fmt.Println(intersect(a, b))
}
