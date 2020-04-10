package main

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/endpoints"
	"github.com/tmornini/udemy-hangman/router"
	"github.com/tmornini/udemy-hangman/secretwords"
)

func main() {
	err := secretwords.Initialize()
	if err != nil {
		panic(err)
	}

	rtr := router.New(
		endpoints.GetRoot{},
		endpoints.GetErrorAnticipated{},
		endpoints.GetErrorUnanticipated{},
	)

	err = http.ListenAndServe(":80", rtr)
	if err != nil {
		panic(err)
	}
}
