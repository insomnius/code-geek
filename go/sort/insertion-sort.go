package main

import (
	"fmt"
	"math/rand"
)

// Insertion sort is a simple sorting algorithm that builds the final sorted array (or list) one item at a time.
// It is much less efficient on large lists than more advanced algorithms such as quicksort, heapsort, or merge sort.
// https://en.wikipedia.org/wiki/Insertion_sort
func insertion_sort() {
	fmt.Println("insertion sort")

	var array_number [50]int
	array_length := len(array_number)

	for i := 0; i < array_length; i++ {
		array_number[i] = rand.Intn(100)
	}

	fmt.Println("random number: ", array_number)

	for i := 0; i < array_length; i++ {
		toswap := i
		lowest := array_number[i]

		for j := i; j >= 0; j-- {

			if array_number[j] > lowest {
				array_number[j+1] = array_number[j]
				toswap = j
			}
		}

		if toswap != i {
			array_number[toswap] = lowest
		}
	}

	fmt.Println("sorted number: ", array_number)
}
