package dto

import (
	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/models"
)

type DifficultyDto struct {
	Id            uuid.UUID `json:"id"`
	Alias         string    `json:"alias"`
	RelativeLevel int       `json:"relativeLevel"`
}

func (dto *DifficultyDto) ToDifficultyEntity() *models.Difficulty {
	return &models.Difficulty{
		Id:            dto.Id,
		Alias:         dto.Alias,
		RelativeLevel: dto.RelativeLevel,
	}
}
