package models

import "github.com/google/uuid"

type Difficulty struct {
	Id            uuid.UUID
	Alias         string
	RelativeLevel int
}
