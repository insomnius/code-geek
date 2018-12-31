package main

import (
	"fmt"
	"strings"
)

/*
You are given a sets of string and your job is to reverse them.

Level: Easy
*/
func reverse_string() {
	fmt.Println("reverse string")

	var to_be_reverse string = "foo bar"
	var reversed_string string

	reversed_string = reverse_string_(to_be_reverse)
	fmt.Println("normal string:", to_be_reverse, ", reversed string:", reversed_string)

	to_be_reverse = "when i was a young boy, my father took me into the city"
	reversed_string = reverse_string_(to_be_reverse)
	fmt.Println("normal string:", to_be_reverse, ", reversed string:", reversed_string)

	to_be_reverse = "and i'm still breathing, i'm still breathing on my own"
	reversed_string = reverse_string_(to_be_reverse)
	fmt.Println("normal string:", to_be_reverse, ", reversed string:", reversed_string)

	to_be_reverse = "maybe i'm not the one that make you smile, but i guarantee i'm the one who make you feel sexy"
	reversed_string = reverse_string_(to_be_reverse)
	fmt.Println("normal string:", to_be_reverse, ", reversed string:", reversed_string)

	to_be_reverse = "lorem ipsum dolor sit amet"
	reversed_string = reverse_string_(to_be_reverse)
	fmt.Println("normal string:", to_be_reverse, ", reversed string:", reversed_string)
}

func reverse_string_(to_be_reverse string) string {
	string_array := strings.Split(to_be_reverse, "")
	string_length := len(string_array)

	start := 0
	last := string_length - 1
	for start < last {
		x := string_array[start]
		string_array[start] = string_array[last]
		string_array[last] = x

		start++
		last--
	}

	return strings.Join(string_array, "")
}
