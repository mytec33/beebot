package wordlegame

import (
	"beebot/scoring"
	"beebot/shared"
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const MAX_GUESSES int = 5

func PlayWordle(wordList []string, pa shared.ProgramArguments) string {
	PrintGameIntro(pa, len(wordList))

	var attempts = 1
	var guess string
	for attempts <= MAX_GUESSES {
		if attempts == 1 {
			guess = pa.StartingWord
		} else {
			guess = ""
		}

		fmt.Printf("Guess %v: %v\n", attempts, guess)

		if guess == pa.Wordle {
			if pa.ResultOnly {
				return fmt.Sprintf("%v tries\n", attempts)
			}
			return fmt.Sprintf("You found the wordle in %v tries!\n", attempts)
		} else if attempts == MAX_GUESSES {
			return fmt.Sprintf("You didn't find the Wordle. The Wordle is: %v\n", pa.Wordle)
		}

		fmt.Printf("Scoring word: %v\n", guess)
		scoring.ScoreWord(guess, pa.Wordle)

		attempts++
	}

	return "Reached unexpected return point"
}

func PrintGameIntro(pa shared.ProgramArguments, wordListCount int) {
	p := message.NewPrinter(language.English)

	p.Printf("Words: %v\n", wordListCount)
	fmt.Printf("First guess: %v\n", pa.StartingWord)
	fmt.Printf("Wordle: %v\n", pa.Wordle)
}
