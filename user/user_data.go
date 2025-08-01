package user

import "github.com/google/uuid"

func GetUserUuid() string {
	id := uuid.New().String()
	return id
}
