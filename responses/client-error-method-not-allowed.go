package responses

import "net/http"

type MethodNotAllowed struct{}

// WriteTo implement interfaces.Responsible
func (res MethodNotAllowed) WriteTo(w http.ResponseWriter) error {
	w.WriteHeader(405)

	return nil
}
