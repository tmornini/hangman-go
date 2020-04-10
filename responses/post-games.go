package entities

import "net/http"

type PostGames GetGame

// WriteResponseTo implement interfaces.Entity
func (ety PostGames) WriteResponseTo(w http.ResponseWriter) error {
	w.Header().Add("Location", "/games/"+ety.ID)
	w.WriteHeader(201)

	return nil
}
