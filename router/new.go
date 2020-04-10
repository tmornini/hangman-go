package router

import "github.com/tmornini/udemy-hangman/interfaces"

func New(endpoints ...interfaces.Endpointable) Server {
	return Server{
		endpoints: endpoints,
	}
}
