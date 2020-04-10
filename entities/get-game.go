package entities

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/bodies"
)

type GetGame bodies.Game

// WriteResponseTo implement interfaces.Entity
func (ety GetGame) WriteResponseTo(w http.ResponseWriter) error {
	w.Header().Add("Location", "/games/"+ety.ID)

	w.WriteHeader(201)

	return nil
}
