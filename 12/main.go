package main

import "fmt"

func main() {
	set := []string{"lol", "kek", "tree", "tree", "go", "kek", "123", "1234"}
	fmt.Println(set[:1])
	// "Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
	// собственное множество."
	// By definition, A is a proper subset of B if all elements of A are elements of B
	// and A doesn't equal B.
	// https://en.wikipedia.org/wiki/Subset
}
