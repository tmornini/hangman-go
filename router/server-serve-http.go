package router

import (
	"fmt"
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
)

func (server Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathResponders := []interfaces.Endpointable{}
	qualifiedResponders := []interfaces.Endpointable{}

	for _, ept := range server.endpoints {
		if ept.RespondsToPathOf(r) {
			pathResponders = append(pathResponders, ept)

			if ept.RespondsToMethodOf(r) {
				qualifiedResponders = append(qualifiedResponders, ept)
			}
		}
	}

	if len(pathResponders) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if len(qualifiedResponders) == 0 {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if len(qualifiedResponders) > 1 {
		panic(
			fmt.Sprintf("multiple endpoints qualify for request: %v", qualifiedResponders),
		)
	}

	dispatch(w, r, qualifiedResponders[0])
}
