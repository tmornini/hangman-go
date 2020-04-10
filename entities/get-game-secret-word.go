package entities

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/bodies"
)

type GetGameSecretWord bodies.GameSecretWord

// WriteResponseTo implement interfaces.Entity
func (ety GetGameSecretWord) WriteResponseTo(w http.ResponseWriter) error {
	w.WriteHeader(200)

	_, err := w.Write([]byte(ety.SecretWord + "\n"))
	if err != nil {
		return err
	}

	return nil
}
