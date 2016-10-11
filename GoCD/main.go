package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Printf("Sum(40, 2) = %d\n", sum(40, 2))
}

func sum(a, b int) int {
	return a + b
}

func even(n int) bool {

	if n%2 == 0 {
		return true
	}
	return false
}
