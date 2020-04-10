package server

import "net/http"

func (server Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, ept := range server.endpoints {
		if ept.RespondsToPathOf(r) {
			dispatch(w, r, ept)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
