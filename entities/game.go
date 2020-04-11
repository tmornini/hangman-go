package entities

type Game struct {
	ID                    string              `json:"-"`
	UnmaskedLetters       []string            `json:"unmasked-letters"`
	WrongGuessesRemaining int                 `json:"wrong-guesses-remaining"`
	GuessedLetterHistory  []GameGuessedLetter `json:"guessed-letter-history"`
}
