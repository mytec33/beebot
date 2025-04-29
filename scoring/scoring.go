package scoring

import (
	"fmt"
	"strings"
)

var Marks [5]int

func GetRequiredLetters(guess string, wordle string) []string {
	var requiredLetters []string
	for x := range 5 {
		if guess[x] == wordle[x] {
			Marks[x] = 2
			requiredLetters = append(requiredLetters, string(guess[x]))
		}
	}

	return requiredLetters
}

func ResetMarks() {
	Marks = [5]int{}
}

func ScoreWord(guess string, wordle string) {
	requiredLetters := GetRequiredLetters(guess, wordle)
	for _, letter := range requiredLetters {
		fmt.Println(letter)
	}

	for x := range 5 {
		if strings.Contains(wordle, string(guess[x])) {
			foundCount := 0
			occurances := 0

		}
	}
}
