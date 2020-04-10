package entities

import "net/http"

type GetErrorAnticipated struct{}

// WriteResponseTo implement interfaces.Entity
func (ety GetErrorAnticipated) WriteResponseTo(w http.ResponseWriter) error {
	w.WriteHeader(503)

	return nil
}
