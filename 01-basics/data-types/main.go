package main

import "fmt"

func main() {
	// Integer types
	var intNumber int = 100
	var intNumber64 int64 = 9000000000

	// Floating point types
	var precise float64 = 19.99999
	var price float32 = 19.99

	// String types
	var name string = "Deepika"

	// Boolean types
	var isStudent bool = true

	// Zero values
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool

	fmt.Println(intNumber)
	fmt.Println(intNumber64)
	fmt.Println(price)
	fmt.Println(precise)
	fmt.Println(name)
	fmt.Println(isStudent)

	fmt.Println(zeroInt)
	fmt.Println(zeroFloat)
	fmt.Println(zeroString)
	fmt.Println(zeroBool)
}
