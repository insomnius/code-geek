package main

import (
	"fmt"
	"math/rand"
)

// Scenario: You are given unsorted array and you have to remove the duplicate values.
// Level: Easy
func remove_duplicate_from_unsorted_array() {
	fmt.Println("remove duplicate from unsorted array")

	var unsorted_array [30]int
	array_length := len(unsorted_array)

	for i := 0; i < array_length; i++ {
		unsorted_array[i] = rand.Intn(10)
	}

	fmt.Println("unsorted array: ", unsorted_array)

	temp := make([]int, array_length)
	j := 0

	for i := 0; i < array_length; i++ {
		if i == 0 {
			temp[j] = unsorted_array[i]
			j++
			continue
		}

		unique := true

		for k := 0; k < i; k++ {
			if unsorted_array[i] == temp[k] {
				unique = false
				break
			}
		}

		if unique {
			temp[j] = unsorted_array[i]
			j++
		}
	}

	result_array := make([]int, j)
	for i := 0; i < j; i++ {
		result_array[i] = temp[i]
	}

	fmt.Println("result array: ", result_array)
}
