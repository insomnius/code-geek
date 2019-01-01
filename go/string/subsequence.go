package main

import (
	"fmt"
)

/*
You are given task to print a possible subsequence of word, so what is subsequence?

In mathematics, a subsequence is a sequence that can be derived from another sequence by deleting some or
no elements without changing the order of the remaining elements. For example, the sequence {A,B,D}
is a subsequence of {A,B,C,D,E,F}  obtained after removal of elements C, E, and F.
The relation of one sequence being the subsequence of another is a preorder.

The subsequence should not be confused with substring {A,B,C,D} which can be derived from the above string {A,B,C,D,E,F}
by deleting substring {E,F}. The substring is a refinement of the subsequence.

The list of all subsequences for the word "apple" would be
"a, ap, al, ae, app, apl, ape, ale, appl, appe, aple, apple, p, pp, pl, pe, ppl, ppe, ple, pple, l, le, e".

source: https://en.wikipedia.org/wiki/Subsequence

Level: Medium
*/
func subsequence() {
	fmt.Println("subsequence")

	var text string
	var empty_string_array []string
	var in_array = map[string]bool{}

	text = "apple"
	fmt.Printf("subsequence of \"%s\" is:\n", text)
	fmt.Printf("%s\n\n", subsequence_(text, len(text), empty_string_array, in_array))
	in_array = map[string]bool{}

	text = "i love you"
	fmt.Printf("subsequence of \"%s\" is:\n", text)
	fmt.Printf("%s\n\n", subsequence_(text, len(text), empty_string_array, in_array))
	in_array = map[string]bool{}

	text = "eat banana"
	fmt.Printf("subsequence of \"%s\" is:\n", text)
	fmt.Printf("%s\n\n", subsequence_(text, len(text), empty_string_array, in_array))
	in_array = map[string]bool{}

	text = "velociraptor eat goat"
	fmt.Printf("subsequence of \"%s\" is:\n", text)
	fmt.Printf("%s\n\n", subsequence_(text, len(text), empty_string_array, in_array))
	in_array = map[string]bool{}
}

func subsequence_(text string, index int, subsequence []string, in_array map[string]bool) []string {
	// If its a blank text or index bigger than text length, then it should be return instead.
	if text == "" || index > len(text) || in_array[text] {
		return subsequence
	}

	// Add text to subsequence array and memoize the text
	subsequence = append(subsequence, text)
	in_array[text] = true

	for i := 1; i <= len(text)-1; i++ {
		for j := i + 2; j <= len(text); j++ {
			combined_text := text[:i] + text[j:]
			if in_array[combined_text] {
				continue
			}
			subsequence = subsequence_(combined_text, len(combined_text), subsequence, in_array)
		}
	}

	for index >= 0 {
		i := index
		index--
		if in_array[text[i:]] == false {
			subsequence = subsequence_(text[i:], index, subsequence, in_array)
		} else if in_array[text[:i]] == false {
			subsequence = subsequence_(text[:i], index, subsequence, in_array)
		}
	}

	return subsequence
}
