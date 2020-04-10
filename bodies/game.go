package bodies

type Game struct {
	ID                    string
	UnmaskedLetters       []string
	WrongGuessesRemaining int
	GuessedLetterHistory  []GameGuessedLetter
}
