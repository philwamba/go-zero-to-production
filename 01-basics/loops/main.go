package main

import "fmt"

func main() {
	// Standard for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// While-style loop
	counter := 1
	for counter <= 5 {
		fmt.Println("While-style loop:", counter)
		counter++
	}

	// Infinite loop with break
	number := 0
	for {
		number++
		if number > 5 {
			break
		}
		fmt.Println("Infinite loop with break:", number)
	}

	// Range loop
	numbers := []int{10, 20, 30}
	for index, value := range numbers {
		fmt.Println("Range loop:", index, value)
	}
}