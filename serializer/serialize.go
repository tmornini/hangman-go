package serializer

import (
	"bytes"
	"encoding/json"

	"github.com/tmornini/udemy-hangman/interfaces"
)

func Serialize(entity interfaces.Responsible, contentType string) ([]byte, error) {
	buffer := &bytes.Buffer{}

	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(entity)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
