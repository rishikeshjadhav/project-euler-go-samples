package main

import "fmt"

func main() {
	fmt.Println("Welcome to Project Euler - Problem 7")
	totalPrimes := 0
	currPrime := 2
	i := 2
	for ; totalPrimes <= 10000; i++ {
		k := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				k++
			}
		}
		if k == 2 {
			currPrime = i
			totalPrimes++
			fmt.Printf("\n%d is prime at position %d", i, totalPrimes)
		} else {
			// fmt.Printf("\n%d is not prime", i)
		}
	}
	fmt.Printf("\nAnswer is %d", currPrime)
}
