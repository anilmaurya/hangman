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
			fmt.Println("Out of changes")
			return
		}
		//evaluate a win!
		if strings.Join(placeholder, "") == word {
			fmt.Println("You WON")
			return
		}
		//Console display

		fmt.Println(placeholder)                // render the placeholder
		fmt.Printf("Chances left: %d", chances) // render the chances left
		fmt.Println("\n")
		for k := range entries {
			fmt.Println(k)
		}
		fmt.Printf("\n Guess a letter or the word: ")

		// take the input
		var str string
		fmt.Scanln(&str)
		if strings.Contains(word, str) {
			for i, v := range word {
				if strings.ContainsRune(str, v) {
					placeholder[i] = str
				}
			}
		} else {
			chances = chances - 1
		}
		entries[str] = true
		// compare and update entries, placeholder and chances.
	}
}
