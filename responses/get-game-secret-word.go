package responses

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/entities"
)

type GetGameSecretWord entities.GameSecretWord

// SerializeTo implement interfaces.Responsible
func (ety GetGameSecretWord) SerializeTo(w http.ResponseWriter) error {
	w.WriteHeader(200)

	_, err := w.Write([]byte(ety.SecretWord + "\n"))
	if err != nil {
		return err
	}

	return nil
}
