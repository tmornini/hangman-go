package responses

import "net/http"

type Unauthorized struct{}

// WriteTo implement interfaces.Responsible
func (res Unauthorized) WriteTo(w http.ResponseWriter) error {
	w.WriteHeader(401)

	return nil
}
