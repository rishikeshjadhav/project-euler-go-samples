package main

import "fmt"

func main() {
	fmt.Println("Welcome to Project Euler - Problem 2")

	first := 0
	second := 1
	fmt.Println(first)
	fmt.Println(second)
	sum := 0
	for i := 0; first < 2178309; i++ {
		first += second
		second += first
		fmt.Println(first)
		fmt.Println(second)
		if first%2 == 0 {
			sum += first
			fmt.Println("Added", first)
		}
		if second%2 == 0 {
			sum += second
			fmt.Println("Added", second)
		}
	}
	fmt.Println(sum)
}
