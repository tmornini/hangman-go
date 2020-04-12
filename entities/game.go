package entities

import "github.com/tmornini/udemy-hangman/interfaces"

type Game struct {
	interfaces.Getable `json:"-"`

	ID                    string              `json:"-"`
	UnmaskedLetters       []string            `json:"unmasked-letters"`
	WrongGuessesRemaining int                 `json:"wrong-guesses-remaining"`
	GuessedLetterHistory  []GameGuessedLetter `json:"guessed-letter-history"`
}
