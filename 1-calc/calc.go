package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("For CLI calculator use this pattern <filename.go> <number1> <operator> <number2>")
		return
	}

	number1, err1 := strconv.ParseFloat(os.Args[1], 64)
	operator := os.Args[2]
	number2, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Both arguments should be numbers")
		return
	}

	switch operator {
	case "+":
		fmt.Printf("Result of the summation: %f + %f = %f\n", number1, number2, number1+number2)
	case "-":
		fmt.Printf("Result of the subtraction: %f - %f = %f\n", number1, number2, number1-number2)
	case "*":
		fmt.Printf("Result of the multiplication: %f * %f = %f\n", number1, number2, number1*number2)
	case "/":
		fmt.Printf("Result of the division: %f / %f = %f\n", number1, number2, number1/number2)
	default:
		fmt.Printf("Unknown operator: %s. Operator could be only +, -, *, /\n", operator)
		return
	}
}
