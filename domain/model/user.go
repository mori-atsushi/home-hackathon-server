package model

import (
	"github.com/google/uuid"
)

type User struct {
	id string
}

func NewUser() User {
	return User{
		id: uuid.Must(uuid.NewRandom()).String(),
	}
}
