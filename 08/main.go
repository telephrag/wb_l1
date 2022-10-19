package main

import (
	"fmt"
)

func setBit(v uint64, pos int, toZero bool) uint64 {
	mask := uint64(1 << (pos)) // Creating a mask of bit we wan't to set
	if toZero {
		return v &^ mask // AND NOT to set masked bit to `0`
	} else {
		return v | mask // OR to set masked bit to `1`
	}
}

func main() {
	var v uint64 = 256 + 128
	fmt.Println(setBit(v, 0, false))
	fmt.Println(setBit(v, 7, true))
}
