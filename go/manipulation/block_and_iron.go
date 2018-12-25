package main

import (
	"fmt"
)

/*
You are a blacksmith and you sell katana complete with its scabbard.
In order to make it you must have one block of wood to make scabbard and one iron in order
to create the complete katana. You must have block of wood before you own iron.
Let's say that block of wood is B and the iron is I, to create one katana you must have:
BI not IB.

To make two katanas, you must have:
BBII or BIBI not BIIB or BIBB

and three katana will need:
BBBIII or
BBIIBI or
BIBBII or
BIBIBI or

Your task is to collect possible values of creating katana.

Level: Hard
*/
func block_and_iron() {
	fmt.Println("block and iron")

	var katana []string
	fmt.Println("2 Katana needs:", block_and_iron_converter("B", "I", "B", "", 2, katana))
	fmt.Println("3 Katana needs:", block_and_iron_converter("B", "I", "B", "", 3, katana[:0]))
	fmt.Println("4 Katana needs:", block_and_iron_converter("B", "I", "B", "", 4, katana[:0]))
	fmt.Println("5 Katana needs:", block_and_iron_converter("B", "I", "B", "", 5, katana[:0]))
}

func block_and_iron_converter(first_word_to_combine, second_word_to_combine, first_word, combined_word string, n int, katana []string) []string {
	for i := 0; i < n; i++ {
		combined_word = combined_word + first_word_to_combine
		word_length := len(combined_word)

		life := 0
		for _, char := range combined_word {
			if string(char) == first_word {
				life++
				continue
			}
			life--
		}

		if life < 0 {
			continue
		}

		if word_length < n*2 {
			if first_word_to_combine == first_word && (life > 1 || life < n) && life < word_length {
				j := word_length - 1
				for i := j; i > 0; i-- {
					if string(combined_word[i]) != first_word {
						break
					}
					j--
				}

				if j != word_length {
					life_check := 0
					for j >= 0 {
						if string(combined_word[j]) == first_word {
							life_check++
							j--
							continue
						}
						life_check--
						j--
					}

					if life_check != 0 {
						continue
					}
				}
			}

			katana = block_and_iron_converter(second_word_to_combine, first_word_to_combine, first_word, combined_word, n, katana)
		} else {
			if word_length > n*2 || life > 0 {
				continue
			}
			katana = append(katana, combined_word)
		}
	}

	return katana
}
