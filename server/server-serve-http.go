package server

import "net/http"

func (server Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, endpoint := range server.endpoints {
		if endpoint.RespondsTo(r) {
			entity, err := endpoint.RespondTo(r)
			if err != nil {
				w.WriteHeader(500)
				return
			}

			entity.WriteResponseTo(w)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
