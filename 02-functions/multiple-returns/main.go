package main

import (
	"errors"
	"fmt"
)

// Function returning multiple values
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Function returning multiple values
func stats(numbers []int) (min int, max int) {
	min = numbers[0]
	max = numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return
}

func main() {
	fmt.Println("Multiple Returns")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	values := []int{1, 2, 3, 4, 5}
	min, max := stats(values)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
