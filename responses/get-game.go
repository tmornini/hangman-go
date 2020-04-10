package responses

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/entities"
)

type GetGame entities.Game

// WriteTo implement interfaces.Responsible
func (ety GetGame) WriteTo(w http.ResponseWriter) error {
	w.Header().Add("Location", "/games/"+ety.ID)

	w.WriteHeader(201)

	return nil
}
