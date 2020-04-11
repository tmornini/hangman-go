package entities

type GameSecretWord struct {
	ID         string `json:"-"`
	SecretWord string `json:"secret-word"`
}
