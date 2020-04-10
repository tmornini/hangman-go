package server

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
)

func Serve(port string, rtr interfaces.Routable) {
	err := http.ListenAndServe(port, rtr)
	if err != nil {
		panic(err)
	}
}
