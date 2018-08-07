package main

import "fmt"

func main() {
	fmt.Println("Welcome to Project Euler - Problem 5")

	for i := 1; ; i++ {
		divisible := true
		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				divisible = false
				break
			}
		}
		if divisible {
			fmt.Println(i)
			break
		}
	}
}
