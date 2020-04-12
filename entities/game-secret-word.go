package entities

import "github.com/tmornini/udemy-hangman/interfaces"

type GameSecretWord struct {
	interfaces.Getable `json:"-"`
	SecretWord         string `json:"secret-word"`
}
