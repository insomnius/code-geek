package main

import (
	"fmt"
	"math/rand"
)

// In computer science, selection sort is a sorting algorithm, specifically an in-place comparison sort.
// It has O(n2) time complexity, making it inefficient on large lists, and generally performs worse than the similar insertion sort.
// Selection sort is noted for its simplicity, and it has performance advantages over more complicated algorithms in certain situations,
// particularly where auxiliary memory is limited.
//
// The algorithm divides the input list into two parts: the sublist of items already sorted,
// which is built up from left to right at the front (left) of the list,
// and the sublist of items remaining to be sorted that occupy the rest of the list.
// Initially, the sorted sublist is empty and the unsorted sublist is the entire input list.
// The algorithm proceeds by finding the smallest (or largest, depending on sorting order) element in the unsorted sublist,
// exchanging (swapping) it with the leftmost unsorted element (putting it in sorted order),
// and moving the sublist boundaries one element to the right.
// https://en.wikipedia.org/wiki/Selection_sort
// Level: Easy
func selection_sort() {
	fmt.Println("selection sort")

	var array_number [50]int
	array_length := len(array_number)

	for i := 0; i < array_length; i++ {
		array_number[i] = rand.Intn(100)
	}

	fmt.Println("random number: ", array_number)

	for i := 0; i < array_length; i++ {
		toswap := i
		lowest := array_number[i]

		for j := i; j < array_length; j++ {

			if array_number[j] < lowest {
				lowest = array_number[j]
				toswap = j
			}
		}

		if toswap != i {
			temp := array_number[i]
			array_number[i] = lowest
			array_number[toswap] = temp
		}
	}

	fmt.Println("sorted number: ", array_number)
}
