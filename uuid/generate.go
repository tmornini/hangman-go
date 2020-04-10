package uuid

import (
	"strings"

	"github.com/google/uuid"
)

func Generate() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
