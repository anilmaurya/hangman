package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "elephant"
	// const to define max const
	const maxChances = 8

	// lookup for entries made by the user.
	entries := make(map[string]bool)

	// list of "_" corrosponding to the number of letters in the word. [ _ _ _ _ _ ]
	placeholder := make([]string, len(word), len(word))
	for i := range placeholder {
		placeholder[i] = "_"
	}

	chances := maxChances
	for {
		//evaluate a loss! If user guesses a wrong letter or the wrong word, they lose a chance.
		if chances == 0 {
			fmt.Println("Out of chances")
			return
		}
		//evaluate a win!
		if strings.Join(placeholder, "") == word {
			fmt.Println("You WON")
			return
		}
		//Console display

		fmt.Printf("placehoder: %v \n", placeholder) // render the placeholder
		fmt.Printf("Chances left: %d", chances)      // render the chances left
		fmt.Println()
		for k := range entries {
			fmt.Printf("%s ", k)
		}
		fmt.Println()
		fmt.Printf("\n Guess a letter or the word: ")

		// take the input
		var str string
		fmt.Scanln(&str)
		if str == word {
			fmt.Println("WON")
			return
		}

		if len(str) == 0 {
			fmt.Println("please enter a character")
			continue
		}

		if len(str) > 1 {
			entries[str] = true
			chances--
			continue
		}

		if !strings.Contains(word, str) {
			_, ok := entries[str]
			if !ok {
				chances--
			}
			entries[str] = true
			continue
		}
		for i, v := range word {
			if strings.ContainsRune(str, v) {
				placeholder[i] = str
			}
		}
		entries[str] = true
		// compare and update entries, placeholder and chances.
	}
}
