package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Project Euler - Problem 3")
	number := 600851475143
	for number != 1 {
		for i := 2; i <= number; i++ {
			if number%i == 0 {
				number = number / i
				fmt.Printf("\n%d - %d\n", i, number)
				break
			}
		}
	}
}
