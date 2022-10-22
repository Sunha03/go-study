package user

import (
	"log"

	"github.com/google/uuid"
)

// StringToUuid : convert(string -> uuid)
func StringToUuid(str string) uuid.UUID {
	u, err := uuid.Parse(str)
	if err != nil {
		log.Fatalln(err)
	}

	return u
}
