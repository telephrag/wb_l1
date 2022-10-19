package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkUnique(str string) bool {
	symbols := make(map[byte]struct{})

	for _, b := range []byte(strings.ToLower(str)) {
		if _, ok := symbols[b]; ok {
			return false
		} else {
			symbols[b] = struct{}{}
		}
	}
	return true
}

func main() {
	r := bufio.NewReader(os.Stdin)

	input, err := r.ReadString('\n')
	if err != nil {
		log.Fatal("string scan failed")
	}

	input = input[:len(input)-1] // removing delimiter (\n) at the end

	fmt.Println(checkUnique(input))
}
