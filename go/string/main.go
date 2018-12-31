package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("\t\tCode Geek\n\n")
	fmt.Println("Language: Go")
	fmt.Println("Category: String")
	hr()

	measure(reverse_string)
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
