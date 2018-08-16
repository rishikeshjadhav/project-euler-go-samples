// Multiples of 3 and 5

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Project Euler - Problem 1")

	sum := 0
	for i := 0; i < 1000; i++ {
		// fmt.Println(i)

		if i%3 == 0 || i%5 == 0 {
			//fmt.Println(i)
			sum += i
		}
	}

	fmt.Println(sum)
}
