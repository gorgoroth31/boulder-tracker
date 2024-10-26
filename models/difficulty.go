package models

import "github.com/google/uuid"

type Difficulty struct {
	Id    uuid.UUID
	alias string
}

func NewDifficulty(id uuid.UUID, alias string) *Difficulty {
	return &Difficulty{Id: id, alias: alias}
}
