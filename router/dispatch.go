package router

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
)

func dispatch(w http.ResponseWriter, r *http.Request, ept interfaces.Endpointable) {
	res, err := ept.Process(r)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	respond(w, res)
}
