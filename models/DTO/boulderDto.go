package dto

import (
	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/models"
)

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

func (boulder *BoulderDto) ToBoulderEntity() *models.Boulder {
	styles := []models.Style{}
	for _, style := range boulder.Style {
		styles = append(styles, *style.ToStyleEntity())
	}
	return &models.Boulder{
		Id:                 boulder.Id,
		ScrewedDifficulty:  *boulder.ScrewedDifficulty.ToDifficultyEntity(),
		FeltLikeDifficulty: *boulder.FeltLikeDifficulty.ToDifficultyEntity(),
		Attempts:           boulder.Attempts,
		SessionsTried:      boulder.SessionsTried,
		Exhausting:         boulder.Exhausting,
		Like:               boulder.Like,
		SessionId:          boulder.SessionId,
		Style:              styles,
	}
}
