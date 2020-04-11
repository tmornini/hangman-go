package responses

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/tmornini/udemy-hangman/entities"
	"github.com/tmornini/udemy-hangman/serializer"
)

type PostGames entities.Game

// WriteTo implement interfaces.Responsible
func (ety PostGames) WriteTo(w http.ResponseWriter) error {
	body, err := serializer.Serialize(ety, "application/json")
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Location", "/games/"+ety.ID)
	w.Header().Add("ETag", fmt.Sprintf("%x", sha256.Sum256(body)))
	w.WriteHeader(201)

	_, err = w.Write(body)
	if err != nil {
		return err
	}

	return nil
}
