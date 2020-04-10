package entities

import "net/http"

type MethodNotAllowed struct{}

// WriteResponseTo implement interfaces.Entity
func (ety MethodNotAllowed) WriteResponseTo(w http.ResponseWriter) error {
	w.WriteHeader(405)

	return nil
}
