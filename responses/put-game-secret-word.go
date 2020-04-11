package responses

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/tmornini/udemy-hangman/entities"
	"github.com/tmornini/udemy-hangman/serializer"
)

type PutGameSecretWord entities.GameSecretWord

// WriteTo implement interfaces.Responsible
func (ety PutGameSecretWord) WriteTo(w http.ResponseWriter) error {
	body, err := serializer.Serialize(ety, "application/json")
	if err != nil {
		return err
	}

	w.Header().Add("ETag", fmt.Sprintf("%x", sha256.Sum256(body)))
	w.WriteHeader(204)

	return nil
}
