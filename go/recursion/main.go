package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("\t\tProgramming Test Interview\n\n")
	fmt.Println("Language: Go")
	fmt.Println("Category: Recursive")
	hr()

	hr()
	measure(block_and_iron)
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
