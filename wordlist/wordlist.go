package wordlist

import (
	"bufio"
	"fmt"
	"os"
)

type WordleLetter struct {
	Letter    string
	Frequency int
}

func (l WordleLetter) String() string {
	return fmt.Sprintf("Letter: %v\tFrequency: %v", l.Letter, l.Frequency)
}

func ReadWordList(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		if len(fileScanner.Text()) == 5 {
			lines = append(lines, fileScanner.Text())
		} else {
			return []string{}, fmt.Errorf("word not equal to 5 characters: %s", fileScanner.Text())
		}
	}
	file.Close()

	return lines, nil
}
