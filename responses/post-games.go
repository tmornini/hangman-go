package responses

import "net/http"

type PostGames GetGame

// WriteTo implement interfaces.Responsible
func (ety PostGames) WriteTo(w http.ResponseWriter) error {
	w.Header().Add("Location", "/games/"+ety.ID)
	w.WriteHeader(201)

	return nil
}
