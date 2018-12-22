package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("\t\tProgramming Test Interview\n\n")
	fmt.Println("Language: Go")
	fmt.Println("Category: Manipulation")
	hr()

	measure(remove_duplicate_from_sorted_array)
	hr()
	measure(remove_duplicate_from_unsorted_array)
	hr()
	measure(reverse_random_array)
	hr()
}

// hr function use to print dividing line
func hr() {
	fmt.Printf("============================================================\n")
}

// measure function use to measure elapsed time to run a function
func measure(f func()) {
	start_time := time.Now()
	f()
	elapsed := float64(time.Since(start_time).Nanoseconds()) / 1000000
	fmt.Println("elapsed time: ", elapsed, "ms")
}
