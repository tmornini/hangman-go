package entities

type GameGuessedLetter struct {
	GuessedLetter string `json:"guessed-letter"`
	UnmaskedCount int    `json:"unmasked-count"`
}
