package wordgenerator

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

func GenerateRandom() (string, error) {
	fd, err := os.Open("/usr/share/dict/words")
	if err != nil {
		return "", err
	}
	defer fd.Close()

	fiveLetterWords := []string{}
	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == 5 {
			fiveLetterWords = append(fiveLetterWords, strings.ToLower(word))
		}
	}

	randomIndex := rand.Int() % len(fiveLetterWords)

	return fiveLetterWords[randomIndex], nil
}
