package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Hangman struct {
	Entries     map[string]bool
	Placeholder []string
	Chances     int
	maxChances  int
	Word        string
	Duration    time.Duration
}

var dev = flag.Bool("dev", false, "dev mode")

func (h *Hangman) setWord() {
	if *dev {
		h.Word = "elephant"
		return
	}
	resp, err := http.Get("https://random-word-api.herokuapp.com/word?number=5")
	if err != nil {
		fmt.Println("server down")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("server down")
	}

	words := []string{}

	err = json.Unmarshal(body, &words)
	if err != nil {
		fmt.Println("error while parsing")
	}

	h.Word = words[0]
}

func (h *Hangman) setPlaceholder() {
	// list of "_" corrosponding to the number of letters in the word. [ _ _ _ _ _ ]
	h.Placeholder = make([]string, len(h.Word), len(h.Word))
	for i := range h.Placeholder {
		h.Placeholder[i] = "_"
	}
}

func (h Hangman) display(timeRemaining float64) {
	fmt.Println()
	fmt.Println()
	fmt.Printf("placehoder: %v \n", h.Placeholder) // render the placeholder
	fmt.Printf("Chances left: %d", h.Chances)      // render the chances left
	fmt.Println()
	fmt.Printf("Entries: ")
	for k := range h.Entries {
		fmt.Printf("%s ", k)
	}
	fmt.Println()
	fmt.Printf("\033[2K\r Time Remaining %v sec", timeRemaining)
	fmt.Printf("\n Guess a letter or the word: ")

}

func (h *Hangman) play(result chan bool, totalTime float64) {
	timeRemaining := totalTime
	start := 0.0
	go func() {
		for range time.Tick(1 * time.Second) {
			start++
			timeRemaining = totalTime - start
		}
	}()
	for {
		//evaluate a loss! If user guesses a wrong letter or the wrong word, they lose a chance.
		if h.Chances == 0 {
			fmt.Println("Out of chances")
			fmt.Println("Correct word is", h.Word)
			result <- false
			return
		}
		//evaluate a win!
		if strings.Join(h.Placeholder, "") == h.Word {
			result <- true
			return
		}
		//Console display

		h.display(timeRemaining)
		// take the input
		var str string
		fmt.Scanln(&str)
		if str == h.Word {
			result <- true
			return
		}

		if len(str) == 0 {
			fmt.Println("please enter a character")
			continue
		}

		if len(str) > 1 {
			h.Entries[str] = true
			h.Chances--
			continue
		}

		if !strings.Contains(h.Word, str) {
			_, ok := h.Entries[str]
			if !ok {
				h.Chances--
			}
			h.Entries[str] = true
			continue
		}
		for i, v := range h.Word {
			if strings.ContainsRune(str, v) {
				h.Placeholder[i] = str
			}
		}
		h.Entries[str] = true
		// compare and update entries, placeholder and chances.
	}
}

func main() {
	flag.Parse()

	h := Hangman{
		Entries:    make(map[string]bool),
		Chances:    8,
		maxChances: 8,
		Duration:   20,
	}
	h.setWord()
	h.setPlaceholder()

	totalTime := h.Duration * time.Second

	t := time.NewTimer(totalTime)

	result := make(chan bool)

	go h.play(result, totalTime.Seconds())

	select {
	case res := <-result:
		if res {
			fmt.Println("WON")
		} else {
			fmt.Println("Loss")
		}

	case <-t.C:
		fmt.Println("Timeout")
		fmt.Println("Correct word is", h.Word)
	}
}
