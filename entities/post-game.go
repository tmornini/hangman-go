package entities

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/bodies"
)

type PostGame bodies.GameGuessedLetter

// WriteResponseTo implement interfaces.Entity
func (ety PostGame) WriteResponseTo(w http.ResponseWriter) error {
	w.WriteHeader(200)

	return nil
}
