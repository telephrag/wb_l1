package main

import (
	"fmt"
)

var justString string

func wrong() { // someFunc()
	// `fmt.Sprint()` convert floats into string only up to certain sign after point
	a := 3.1415926535897932384626433832795028841971693993751058209
	a += 0.0000000000000004 // signs past ...932 exist but they are omitted by `fmt.Sprint()`
	v := fmt.Sprint(3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470)
	// fmt.Println(v, a)    // 3.141592653589793 3.1415926535897936
	justString = v[:100] // out of bounds panic due to reasons above
}

func correct() {
	v := fmt.Sprint(3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470)
	if len(v) >= 100 { // length check fixes the problem
		justString = v[:100]
	} else {
		justString = v[:]
	}
}

func main() {
	correct()
	// wrong() // will panic
	fmt.Println(justString)
}
