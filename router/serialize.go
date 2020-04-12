package router

import (
	"log"
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
)

func respond(w http.ResponseWriter, res interfaces.Responsible) {
	err := res.SerializeTo(w)
	if err != nil {
		log.Println(err)
	}
}
