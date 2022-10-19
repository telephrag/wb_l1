package main

import (
	"errors"
	"fmt"
)

func deleteAt(s *[]int, pos int) (self *[]int, err error) {
	if pos >= len(*s) {
		return s, errors.New("position is out of bounds")
	}

	// Shifting to the left elements right of `pos` saves us reallocation that would have occured
	// if we used `*s = append((*s)[:pos], (*s)[pos+1:])`, minus the bounds check.
	for i := pos; i < len(*s)-1; i++ {
		(*s)[i] = (*s)[i+1]
	}

	*s = (*s)[:len(*s)-1]

	return s, nil
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println(deleteAt(&s, 7))
}
