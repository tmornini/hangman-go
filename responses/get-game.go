package responses

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/entities"
)

type GetGame entities.Game

// SerializeTo implement interfaces.Responsible
func (ety GetGame) SerializeTo(w http.ResponseWriter) error {
	w.Header().Add("Location", "/games/"+ety.ID)

	w.WriteHeader(201)

	return nil
}
