package dto

import "github.com/google/uuid"

type StyleDto struct {
	Id    uuid.UUID `json:"id"`
	Alias string    `json:"alias"`
}
