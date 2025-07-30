package models

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID
	UserName  string
	Email     string
	IsDeleted bool
	Sessions  []Session
}
