package dto

import "github.com/google/uuid"

type BoulderDto struct {
	Id                 uuid.UUID     `json:"id"`
	ScrewedDifficulty  DifficultyDto `json:"screwedDifficulty"`
	FeltLikeDifficulty DifficultyDto `json:"feltLikeDifficulty"`
	Attempts           int           `json:"attempts"`
	SessionsTried      int           `json:"sessionsTried"`
	Exhausting         bool          `json:"exhausting"`
	Style              []StyleDto    `json:"style"`
	Like               bool          `json:"like"`
	SessionId          uuid.UUID     `json:"SessionId"`
}
