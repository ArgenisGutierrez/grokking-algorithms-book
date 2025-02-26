// Package main recursive
package main

import "fmt"

func countdown(i int) {
	fmt.Println(i)
	// base case
	if i <= 0 {
		return
	} else { //recursive
		countdown(i - 1)
	}
}

func factorial(i int) int {
	if i == 1 {
		return 1
	} else {
		return i * factorial(i-1)
	}
}

func main() {
	countdown(5)
	fmt.Println(factorial(5))
}
