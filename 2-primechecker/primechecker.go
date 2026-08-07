package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("For primechecker.go you need to use this pattern: <filename.go> <number>")
		return
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Your argument should be a number")
	}

	if isPrime(number) {
		fmt.Println("The number is prime")
	} else {
		fmt.Println("The number is not prime")
	}
}
