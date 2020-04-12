package responses

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/entities"
	"github.com/tmornini/udemy-hangman/interfaces"
)

type PostGame struct {
	interfaces.Getable `json:"-"`

	entities.GameGuessedLetter
}

// SerializeTo implement interfaces.Responsible
func (ety PostGame) SerializeTo(w http.ResponseWriter) error {
	w.WriteHeader(200)

	return nil
}
