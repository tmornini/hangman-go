package router

import (
	"fmt"
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
)

func (rtr Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathResponders := []interfaces.Endpointable{}
	qualifiedResponders := []interfaces.Endpointable{}

	for _, ept := range rtr.endpoints {
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

	ept := qualifiedResponders[0]

	res := ept.Authorize(r)
	if res != nil {
		respond(w, res)
	} else {
		dispatch(w, r, ept)
	}
}
