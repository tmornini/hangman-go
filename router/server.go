package router

import "github.com/tmornini/udemy-hangman/interfaces"

type Server struct {
	endpoints []interfaces.Endpointable
}
