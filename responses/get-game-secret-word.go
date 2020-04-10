package responses

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/entities"
)

type GetGameSecretWord entities.GameSecretWord

// WriteTo implement interfaces.Responsible
func (ety GetGameSecretWord) WriteTo(w http.ResponseWriter) error {
	w.WriteHeader(200)

	_, err := w.Write([]byte(ety.SecretWord + "\n"))
	if err != nil {
		return err
	}

	return nil
}
