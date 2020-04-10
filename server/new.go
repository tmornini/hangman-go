package server

import "github.com/tmornini/udemy-hangman/interfaces"

func New(endpoints ...interfaces.Endpoint) Server {
	return Server{
		endpoints: endpoints,
	}
}
