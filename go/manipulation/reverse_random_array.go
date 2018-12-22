package main

import (
	"fmt"
	"math/rand"
)

// Scenario, you are given a random array and you have to reverse its order.
// Level: Easy
func reverse_random_array() {
	fmt.Println("reverse random array")

	var random_array [15]int
	array_length := len(random_array)

	for i := 0; i < array_length; i++ {
		random_array[i] = rand.Intn(100)
	}

	fmt.Println("random array: ", random_array)

	for i := 0; i < array_length; i++ {
		temp := random_array[i]
		random_array[i] = random_array[array_length-1]
		random_array[array_length-1] = temp
		array_length--
	}

	fmt.Println("reversed_array: ", random_array)
}
