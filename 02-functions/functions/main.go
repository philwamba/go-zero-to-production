package main

import "fmt"

// Simple function
func greet(name string) string {
	return "Hello, " + name
}

// Function with parameters and return values
func add(a int, b int) int {
	return a + b
}

// Fuction with named return values
func multiply(a int, b int) (result int) {
	result = a * b
	return
}

func main() {
	message := greet("Deepika")
	sum := add(5, 10)
	product := multiply(4, 5)

	fmt.Println(message)
	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)
}
