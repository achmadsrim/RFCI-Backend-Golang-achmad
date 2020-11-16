package main

import (
	"fmt"
)

func main() {
	collection := []int{3, 5, 2, -4, 8, 11}
	fmt.Println(foo(7, collection))
}

func foo(input int, numbers []int) (int, []int) {
	var result []int
	var nextIndex int

	for idx, number := range numbers {
		nextIndex = idx + 1

		if nextIndex >= len(numbers) {
			break
		}
		nextNumber := numbers[nextIndex]
		if number+nextNumber == input {
			result = append(result, number)
			result = append(result, nextNumber)
		}
	}
	return input, result
}
