package main

import (
	"fmt"
	"math/rand"
)

// Simple trick to do array sort in just one loop
func sort_in_one_loop() {
	fmt.Println("sort in one loop")

	var array_number [50]int
	array_length := len(array_number)

	for i := 0; i < array_length; i++ {
		array_number[i] = rand.Intn(100)
	}

	fmt.Println("random number: ", array_number)

	i := 0
	for i < array_length {

		if i != array_length-1 &&
			i >= 0 &&
			array_number[i] > array_number[i+1] {
			x := array_number[i]
			array_number[i] = array_number[i+1]
			array_number[i+1] = x
			i -= 2
		}

		i++
	}

	fmt.Println("sorted number: ", array_number)
}
