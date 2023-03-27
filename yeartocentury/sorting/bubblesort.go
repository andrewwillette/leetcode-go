package sorting

import "fmt"

func Bubblesort(input []int) []int {
	for i := len(input); i > 0; i-- {
		for j := 1; j < i; j++ {
			if input[j-1] > input[j] {
				intermediate := input[j]
				input[j] = input[j-1]
				input[j-1] = intermediate
			}
			fmt.Printf("insidej %v\n", input)
		}
		fmt.Printf("outsdej %v\n", input)
	}
	return input
}
