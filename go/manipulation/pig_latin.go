package main

import "fmt"

/**
Pig Latin is an argot ("secret" language) in which word in english are altered to
conceal the words from others not familiar with the rules.

You convert a word Pig Latin by transfering the initial consonant or consonant cluster
of the word to the end of the word with -ay appended to the end. if the word start with
vowel (including y), append -yay to the end. For example, "pig" would become igpay
(taking the 'p' and 'ay' to create a suffix)

Examples:
Input: say
Output: aysay

Input: english
Output: englishyay

Input: smile
Output: ilesmay

Level: Easy
*/
func pig_latin() {
	fmt.Println("pig latin")

	var bag_of_words = []string{
		"when", "i", "was", "a", "young", "boy", "my", "father", "took", "me", "into", "city", "to", "see",
		"a", "marching", "band",
	}
	bag_of_words_length := len(bag_of_words)
	fmt.Println("bag of words:", bag_of_words)

	for i := 0; i < bag_of_words_length; i++ {
		fmt.Printf(`english word: "%s", `, bag_of_words[i])
		bag_of_words[i] = pig_latin_converter(bag_of_words[i])
		fmt.Printf(`pig latin word: "%s"%s`, bag_of_words[i], "\n")
	}

	fmt.Println("converted words:", bag_of_words)
}

func pig_latin_converter(word string) (pig_latin_word string) {
	vowel_plus_y := "aioeoy"

	for _, vowel := range vowel_plus_y {
		if string(word[0]) == string(vowel) {
			return word + "ay"
		}
	}

	first_char := string(word[0])
	new_word := word[1:] + first_char
	return pig_latin_converter(new_word)
}
