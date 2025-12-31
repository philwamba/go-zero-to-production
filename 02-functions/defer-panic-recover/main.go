package main

import "fmt"

func safeOperation() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	fmt.Println("Starting operation")

	// Simulate unexpected error
	panic("Something went wrong!")

	// This line will not execute
	fmt.Println("Operation completed successfully")
}

func fileSimulation() {
	defer fmt.Println("File closed")

	fmt.Println("File opened")
}

func main() {
	fileSimulation()

	fmt.Println("Calling safe operation")
	safeOperation()

	fmt.Println("Program continues normally")
}
