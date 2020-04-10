package main

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/endpoints"
	"github.com/tmornini/udemy-hangman/server"
)

func main() {
	svr := server.New(
		endpoints.Root{},
		endpoints.ErrorAnticipated{},
		endpoints.ErrorUnanticipated{},
	)

	err := http.ListenAndServe(":80", svr)
	if err != nil {
		panic(err)
	}
}
