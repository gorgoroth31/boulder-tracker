package mapper

import (
	"github.com/gorgoroth31/boulder-tracker/models"
	"github.com/gorgoroth31/boulder-tracker/models/dto"
)

func ToDifficultyDto(entity *models.Difficulty) *dto.DifficultyDto {
	return &dto.DifficultyDto{
		Id:            entity.Id,
		Alias:         entity.Alias,
		RelativeLevel: entity.RelativeLevel,
	}
}
