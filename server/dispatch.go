package server

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
)

func dispatch(w http.ResponseWriter, r *http.Request, ept interfaces.Endpoint) {
	ety, err := ept.Validate(r)
	if ety != nil {
		ety.WriteResponseTo(w)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		return
	}

	ety, err = ept.RespondTo(r)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	ety.WriteResponseTo(w)
}
