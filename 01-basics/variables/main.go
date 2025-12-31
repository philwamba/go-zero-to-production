package main

import "fmt"

func main() {
	// Declare a variable
	var name string = "Deepika"
	var age int = 25
	var isStudent bool = true
	var height float64 = 1.65

	// Multiple variables
	var (
		city string = "New York"
		isFemale bool = true
	)

	// Short variable declaration
	country := "USA"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isStudent)
	fmt.Println(height)
	fmt.Println(city)
	fmt.Println(country)
	fmt.Println(isFemale)
}
