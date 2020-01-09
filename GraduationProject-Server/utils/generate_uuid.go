package utils

import (
	"github.com/gofrs/uuid"
	"log"
)

// GenerateUUID 返回一个唯一的UUID
func GenerateUUID() string {
	u, err := uuid.NewV4()
	if err != nil {
		log.Println("utils.GenerateUUID:", err)
	}
	return u.String()
}
