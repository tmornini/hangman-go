package secretwords

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var rng *rand.Rand
var words []string

func Initialize() error {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))

	wordFilename := os.Getenv("WORD_FILE_PATHNAME")
	if wordFilename == "" {
		wordFilename = "/usr/share/dict/words"
	}

	wordReader, err := os.Open(wordFilename)
	if err != nil {
		return err
	}

	wordScanner := bufio.NewScanner(wordReader)

	for wordScanner.Scan() {
		words = append(words, wordScanner.Text())
	}

	return nil
}

func ChooseOne() string {
	return words[rng.Intn(len(words))]
}
