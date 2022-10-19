package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func reverseWords(words []string) []string {
	l := len(words)

	res := make([]string, l)
	for i, w := range words {
		res[l-i-1] = w
	}
	return res
}

func main() {
	r := bufio.NewReader(os.Stdin)

	input, err := r.ReadString('\n')
	if err != nil {
		log.Fatal("string scan failed")
	}

	input = input[:len(input)-1] // removing delimiter (\n) at the end

	words := strings.Split(input, " ")

	fmt.Println(reverseWords(words))
}
