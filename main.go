package main

import (
	"beebot/shared"
	"beebot/wordlegame"
	"beebot/wordlist"
	"errors"
	"flag"
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stderr, nil))

func main() {
	pa, err := ParseArgs()
	if err != nil {
		logger.Error("invalid program args", "message", err.Error())
		flag.Usage()

		os.Exit(1)
	}

	wordList, err := wordlist.ReadWordList(pa.WordListFile)
	if err != nil {
		logger.Error("unable to read word list file", "message", err.Error())
		os.Exit(2)
	}

	wordlegame.PlayWordle(wordList, pa)

}

func ParseArgs() (shared.ProgramArguments, error) {
	var resultOnly = flag.Bool("result-only", false, "show only the result, little detail")
	var startingWord = flag.String("starting-word", "", "your starting word")
	var wordle = flag.String("wordle", "", "wordle word of the day")
	var wordListFile = flag.String("word-list-file", "", "word list to play from")
	flag.Parse()

	if len(*startingWord) != 5 {
		return shared.ProgramArguments{}, errors.New("invalid argument --starting-word must be exactly five characters")
	}

	if len(*wordle) != 5 {
		return shared.ProgramArguments{}, errors.New("invalid argument --wordle must be exactly five characters")
	}

	return shared.ProgramArguments{
		ResultOnly:   *resultOnly,
		StartingWord: *startingWord,
		Wordle:       *wordle,
		WordListFile: *wordListFile,
	}, nil
}
