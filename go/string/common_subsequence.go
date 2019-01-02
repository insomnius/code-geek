package main

import "fmt"

/*
The longest common subsequence (LCS) problem is the problem of finding the longest subsequence common
to all sequences in a set of sequences (often just two sequences).

source: https://en.wikipedia.org/wiki/Longest_common_subsequence_problem

You are given 2 string and you should return the longest common subsequence word of it.
Level: Medium
*/
func common_subsequence() {
	fmt.Println("common subsequence")

	var string_1 string
	var string_2 string
	var memoize map[int]map[int]string

	string_1 = "VAIXBKNDFLDMORKQUBJP"
	string_2 = "XGRKPROQISGKLNMWPUQO"
	memoize = map[int]map[int]string{}
	fmt.Printf("first string: %s, second string: %s.\nThe common subsequence is: %s\n\n",
		string_1, string_2, common_subsequence_(string_1, string_2, 0, 0, memoize))

	string_1 = "apple"
	string_2 = "banana"
	memoize = map[int]map[int]string{}
	fmt.Printf("first string: %s, second string: %s.\nThe common subsequence is: %s\n\n",
		string_1, string_2, common_subsequence_(string_1, string_2, 0, 0, memoize))

	string_1 = "velociraptor"
	string_2 = "marsupilami"
	memoize = map[int]map[int]string{}
	fmt.Printf("first string: %s, second string: %s.\nThe common subsequence is: %s\n\n",
		string_1, string_2, common_subsequence_(string_1, string_2, 0, 0, memoize))

	string_1 = "i love you"
	string_2 = "i hate you"
	memoize = map[int]map[int]string{}
	fmt.Printf("first string: %s, second string: %s.\nThe common subsequence is: %s\n\n",
		string_1, string_2, common_subsequence_(string_1, string_2, 0, 0, memoize))
}

func common_subsequence_(s_1, s_2 string, x, y int, memoize map[int]map[int]string) (result string) {
	if len(memoize[x][y]) > 0 {
		return memoize[x][y]
	}

	if x == len(s_1) || y == len(s_2) {
		return ""
	} else if s_1[x] == s_2[y] {
		if len(memoize[x]) == 0 {
			memoize[x] = map[int]string{}
		}
		tmp := result + string(s_1[x]) + common_subsequence_(s_1, s_2, x+1, y+1, memoize)
		memoize[x][y] = tmp
		return tmp
	} else if s_1[x] != s_2[y] {
		r1 := result + common_subsequence_(s_1, s_2, x, y+1, memoize)
		r2 := result + common_subsequence_(s_1, s_2, x+1, y, memoize)
		if len(r1) > len(r2) {
			result = r1
		} else {
			result = r2
		}

		if len(memoize[x]) == 0 {
			memoize[x] = map[int]string{}
		}
		memoize[x][y] = result
		return result
	}

	return result
}
