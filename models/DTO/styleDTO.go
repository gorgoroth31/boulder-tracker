package dto

import (
	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/models"
)

type StyleDto struct {
	Id    uuid.UUID `json:"id"`
	Alias string    `json:"alias"`
}

func (dto *StyleDto) ToStyleEntity() *models.Style {
	return &models.Style{
		Id:    dto.Id,
		Alias: dto.Alias,
	}
}
