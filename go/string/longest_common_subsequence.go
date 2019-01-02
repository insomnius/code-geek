package main

import "fmt"

/*
	You are given 2 string and you should return the longest common subsequence
	Level: Medium
*/
func longest_common_subsequence() {
	fmt.Println("longest common subsequence")

	var string_1 string
	var string_2 string
	var memoize = map[int]map[int]int{}

	string_1 = "banana"
	string_2 = "ohana"
	memoize = map[int]map[int]int{}
	fmt.Printf("first string: %s, second string: %s.\nThe longest common subsequence length is: %d\n\n",
		string_1, string_2, lcs(string_1, string_2, len(string_1), len(string_2), memoize))

	string_1 = "apple"
	string_2 = "apple"
	memoize = map[int]map[int]int{}
	fmt.Printf("first string: %s, second string: %s.\nThe longest common subsequence length is: %d\n\n",
		string_1, string_2, lcs(string_1, string_2, len(string_1), len(string_2), memoize))

	string_1 = "velociraptor"
	string_2 = "marsupilami"
	memoize = map[int]map[int]int{}
	fmt.Printf("first string: %s, second string: %s.\nThe longest common subsequence length is: %d\n\n",
		string_1, string_2, lcs(string_1, string_2, len(string_1), len(string_2), memoize))

	string_1 = "i love you"
	string_2 = "i hate you"
	memoize = map[int]map[int]int{}
	fmt.Printf("first string: %s, second string: %s.\nThe longest common subsequence length is: %d\n\n",
		string_1, string_2, lcs(string_1, string_2, len(string_1), len(string_2), memoize))

}

func lcs(x, y string, i, j int, memoize map[int]map[int]int) int {
	if memoize[i][j] != 0 {
		return memoize[i][j]
	}

	result := 1

	if i == 0 || j == 0 {
		return 0
	} else if x[i-1] == y[j-1] {
		result = 1 + lcs(x, y, i-1, j-1, memoize)
	} else if x[i-1] != y[j-1] {
		res1 := lcs(x, y, i-1, j, memoize)
		res2 := lcs(x, y, i, j-1, memoize)

		if res1 > res2 {
			result = res1
		} else {
			result = res2
		}
	}

	if len(memoize[i]) == 0 {
		memoize[i] = map[int]int{}
	}
	memoize[i][j] = result
	return result
}
