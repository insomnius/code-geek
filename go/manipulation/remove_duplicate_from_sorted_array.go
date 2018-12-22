package main

import (
	"fmt"
)

// Scenario: You are given sorted array and you have to remove all of duplicate values.
// Level: Easy
func remove_duplicate_from_sorted_array() {
	fmt.Println("remove duplicate from sorted array")

	var array_number = [20]int{1, 1, 2, 2, 2, 3, 3, 3, 4, 5, 6, 6, 6, 7, 7, 7, 7, 8, 8, 9}
	array_length := len(array_number)
	var temp [20]int

	fmt.Println("sorted random number: ", array_number)

	j := 0

	for i := 0; i < array_length-1; i++ {
		if array_number[i] != array_number[i+1] {
			temp[j] = array_number[i]
			j++
		}
	}

	temp[j] = array_number[array_length-1]
	j++
	result_array := make([]int, j)

	for i := 0; i < j; i++ {
		result_array[i] = temp[i]
	}

	fmt.Println("result array: ", result_array)
}
