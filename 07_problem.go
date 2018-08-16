package main

import "fmt"

var totalPrimes int = 0
var currPrime int = 2

func checkIfPrime(i int) {
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

func main() {
	fmt.Println("Welcome to Project Euler - Problem 7")
	for i := 2; totalPrimes <= 100; i++ {
		checkIfPrime(i)
	}
	fmt.Printf("\nAnswer is %d", currPrime)
}
