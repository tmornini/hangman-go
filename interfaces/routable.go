package interfaces

import "net/http"

type Routable interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
