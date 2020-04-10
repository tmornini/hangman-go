package responses

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/entities"
)

type PostGame entities.GameGuessedLetter

// WriteTo implement interfaces.Responsible
func (ety PostGame) WriteTo(w http.ResponseWriter) error {
	w.WriteHeader(200)

	return nil
}
