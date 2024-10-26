package dto

import "github.com/google/uuid"

type DifficultyDto struct {
	Id    uuid.UUID `json:"id"`
	Alias string    `json:"alias"`
}
