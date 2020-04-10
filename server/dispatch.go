package server

import (
	"log"
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
)

func dispatch(w http.ResponseWriter, r *http.Request, ept interfaces.Endpoint) {
	ety, err := ept.Validate(r)
	if ety != nil {
		err = ety.WriteResponseTo(w)
		if err != nil {
			log.Println(err)
		}

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

	err = ety.WriteResponseTo(w)
	if err != nil {
		log.Println(err)
	}
}
