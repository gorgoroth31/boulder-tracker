package dto

import "github.com/google/uuid"

type DifficultyDto struct {
	Id    uuid.UUID
	Alias string
}
