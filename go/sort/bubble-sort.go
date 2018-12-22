package main

import (
	"fmt"
	"math/rand"
)

// Bubble sort, sometimes referred to as sinking sort, is a simple sorting algorithm that repeatedly steps through the list,
// compares adjacent pairs and swaps them if they are in the wrong order. The pass through the list is repeated until the list is sorted.
// The algorithm, which is a comparison sort, is named for the way smaller or larger elements "bubble" to the top of the list.
// Although the algorithm is simple, it is too slow and impractical for most problems even when compared to insertion sort.
// Bubble sort can be practical if the input is in mostly sorted order with some out-of-order elements nearly in position.
// https://en.wikipedia.org/wiki/Bubble_sort
// Level: Easy
func bubble_sort() {
	fmt.Println("bubble sort")

	var array_number [50]int
	array_length := len(array_number)

	for i := 0; i < array_length; i++ {
		array_number[i] = rand.Intn(100)
	}

	fmt.Println("random number: ", array_number)

	for array_length > 0 {
		new_length := 0

		for i := 1; i < array_length; i++ {
			x := array_number[i-1]
			y := array_number[i]

			if x > y {
				array_number[i-1] = y
				array_number[i] = x
				new_length = i
			}
		}

		array_length = new_length
	}

	fmt.Println("sorted number: ", array_number)
}
