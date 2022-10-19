package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func reverseStr(str string) string {
	res := make([]byte, len(str))
	for i, b := range []byte(str) {
		res[len(res)-i-1] = b
	}
	return string(res)
}

func main() {
	r := bufio.NewReader(os.Stdin)

	input, err := r.ReadString('\n')
	if err != nil {
		log.Fatal("string scan failed")
	}

	input = input[:len(input)-1] // removing delimiter (\n) at the end

	fmt.Println(reverseStr(input))
}
