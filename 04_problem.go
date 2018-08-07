package main

import (
	"fmt"
	"strconv"
)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func checkIfPalindrome(currentString string) bool {
	isPalindrome := false
	firstPart := currentString[0:(len(currentString) / 2)]
	secondPart := currentString[(len(currentString) / 2):len(currentString)]
	if len(currentString)%2 != 0 {
		secondPart = secondPart[1:len(secondPart)]
	}
	if firstPart == Reverse(secondPart) {
		isPalindrome = true
	}
	return isPalindrome
}

func main() {
	fmt.Println("Welcome to Project Euler - Problem 3")
	largestPalindrome := 0
	for a := 999; a > 99; a-- {
		for b := 999; b > 99; b-- {
			numProduct := a * b
			if checkIfPalindrome(strconv.Itoa(numProduct)) {
				if largestPalindrome < numProduct {
					largestPalindrome = numProduct
					fmt.Printf("\n%d - %d and %d\n", largestPalindrome, a, b)
				}
			}
		}
	}
	fmt.Println(largestPalindrome)
}
