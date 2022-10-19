package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {

	// Examples:
	// 49612349876129846198 * 461324861293481 = 22887410425049250569240191870035238
	// 49612349876129846198 / 461324861293481 = 107543.1957797044923564
	// --------------------------------------------------------
	// Values match those ouput by calculator.net
	fmt.Println("Input two numbers:")

	var as, bs string
	var a, b decimal.Decimal
	for {
		fmt.Scanf("%s %s", &as, &bs)

		var err error
		a, err = decimal.NewFromString(as)
		if err != nil {
			fmt.Println("Invalid value of the first number.")
			continue
		}
		b, err = decimal.NewFromString(bs)
		if err != nil {
			fmt.Println("Invalid value of the second number.")
			continue
		}

		break
	}

	var op string
	fmt.Println("Choose operation from the list: +, -, *, /")
	for {
		fmt.Scanf("%s", &op)

		switch {
		case op == "+":
			fmt.Println(a.Add(b))
			return
		case op == "-":
			fmt.Println(a.Sub(b))
			return
		case op == "*":
			fmt.Println(a.Mul(b))
			return
		case op == "/":
			fmt.Println(a.Div(b))
			return
		default:
			fmt.Printf("\"%s\" is not a valid operand.\n", op)
		}
	}
}
