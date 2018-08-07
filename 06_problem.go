package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Project Euler - Problem 6")

	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += (i * i)
	}
	fmt.Println(sum1)
	sum2 := 0
	for i := 0; i <= 100; i++ {
		sum2 += i
	}
	sum2 *= sum2
	fmt.Println(sum2)
	fmt.Println(sum2 - sum1)
}
