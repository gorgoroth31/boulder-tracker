package models

import (
	"github.com/google/uuid"
)

type Style struct {
	Id    uuid.UUID
	Alias string
}

func NewStyle(id uuid.UUID, alias string) *Style {
	return &Style{Id: id, Alias: alias}
}
