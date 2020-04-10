package entities

type Game struct {
	ID                    string
	UnmaskedLetters       []string
	WrongGuessesRemaining int
	GuessedLetterHistory  []GameGuessedLetter
}
