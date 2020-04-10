package main

import (
	"github.com/tmornini/udemy-hangman/endpoints"
	"github.com/tmornini/udemy-hangman/router"
	"github.com/tmornini/udemy-hangman/secretwords"
	"github.com/tmornini/udemy-hangman/server"
)

func main() {
	err := secretwords.Initialize()
	if err != nil {
		panic(err)
	}

	server.Serve(
		":80",
		router.New(
			endpoints.GetErrorAnticipated{},
			endpoints.GetErrorUnanticipated{},
		),
	)
}
