package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	PrivateId  uuid.UUID `json:"privateId"`
	PublicId   uuid.UUID `json:"publicId"`
	LastActive time.Time `json:"lastActive"`
}

func NewUser(privateIdString string) *User {
	return &User{
		uuid.MustParse(privateIdString),
		uuid.New(),
		time.Now(),
	}
}

func (user *User) String() string {
	privateId := user.PrivateId.String()
	publicId := user.PublicId.String()
	lastActive := user.LastActive.UTC().Format("02.01.2006 15:04:05 MTS")
	return fmt.Sprintf("{ privateId: %s , publicId: %s , lastActive: %s }", privateId, publicId, lastActive)
}
