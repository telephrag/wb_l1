package main

import (
	"fmt"
	"reflect"
)

func getTypes(input []interface{}) []reflect.Type {
	res := []reflect.Type{}
	for _, elem := range input {
		res = append(res, reflect.TypeOf(elem))
	}
	return res
}

func main() {
	input := []interface{}{123, "str", false, make(chan struct{}, 1)}

	fmt.Println(getTypes(input))
}
