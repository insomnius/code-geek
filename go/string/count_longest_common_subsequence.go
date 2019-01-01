package main

import "fmt"

func count_longest_common_subsequence() {
	fmt.Println("count_longest_common_subsequence")

	var string_1 string
	var string_2 string

	string_1 = "banana"
	string_2 = "ohana"
	fmt.Printf("first string: %s, second string: %s.\nThe longest common subsequence length is: %d\n\n",
		string_1, string_2, clcs(string_1, string_2, len(string_1), len(string_2)))

	string_1 = "apple"
	string_2 = "apple"
	fmt.Printf("first string: %s, second string: %s.\nThe longest common subsequence length is: %d\n\n",
		string_1, string_2, clcs(string_1, string_2, len(string_1), len(string_2)))

	string_1 = "velociraptor"
	string_2 = "marsupilami"
	fmt.Printf("first string: %s, second string: %s.\nThe longest common subsequence length is: %d\n\n",
		string_1, string_2, clcs(string_1, string_2, len(string_1)-1, len(string_2)-1))

	string_1 = "i love you"
	string_2 = "i hate you"
	fmt.Printf("first string: %s, second string: %s.\nThe longest common subsequence length is: %d\n\n",
		string_1, string_2, clcs(string_1, string_2, len(string_1)-1, len(string_2)-1))
}

func clcs(x, y string, i, j int) int {
	result := 1

	if i == 0 || j == 0 {
		return 0
	} else if x[i-1] == y[j-1] {
		result = 1 + clcs(x, y, i-1, j-1)
	} else if x[i-1] != y[j-1] {
		res1 := clcs(x, y, i-1, j)
		res2 := clcs(x, y, i, j-1)

		if res1 > res2 {
			result = res1
		} else {
			result = res2
		}
	}

	return result
}
