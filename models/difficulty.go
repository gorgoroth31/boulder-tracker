package models

import (
	"github.com/google/uuid"
)

type Difficulty struct {
	Id            uuid.UUID
	Alias         string
	RelativeLevel int
}

func NewDifficulty(Id uuid.UUID, Alias string, RelativeLevel int) *Difficulty {
	return &Difficulty{
		Id:            Id,
		Alias:         Alias,
		RelativeLevel: RelativeLevel,
	}
}
