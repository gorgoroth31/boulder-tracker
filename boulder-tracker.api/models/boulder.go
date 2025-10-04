package models

import (
	"github.com/google/uuid"
)

type Boulder struct {
	Id                   uuid.UUID
	ScrewedDifficulty    Difficulty
	FeltLikeDifficulty   Difficulty
	Attempts             int
	SessionsTried        int
	Exhausting           bool
	Style                []Style
	Like                 bool
	SessionId            uuid.UUID
	ScrewedDifficultyId  uuid.UUID
	FeltLikeDifficultyId uuid.UUID
}

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

func NewBoulder(Id uuid.UUID, screwedDifficulty Difficulty,
	feltLikeDifficulty Difficulty,
	attempts int,
	sessionsTried int,
	exhausting bool,
	style []Style, like bool) *Boulder {
	return &Boulder{
		Id:                 Id,
		ScrewedDifficulty:  screwedDifficulty,
		FeltLikeDifficulty: feltLikeDifficulty,
		Attempts:           attempts,
		SessionsTried:      sessionsTried,
		Exhausting:         exhausting,
		Style:              style,
		Like:               like,
	}
}

func (boulder *BoulderDto) ToBoulderEntity() *Boulder {
	styles := []Style{}
	for _, style := range boulder.Style {
		styles = append(styles, *style.ToStyleEntity())
	}
	return &Boulder{
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

func (boulder *Boulder) ToBoulderDTO() *BoulderDto {
	styles := []StyleDto{}
	for _, style := range boulder.Style {
		styles = append(styles, *style.ToStyleDTO())
	}

	return &BoulderDto{
		Id:                 boulder.Id,
		ScrewedDifficulty:  *boulder.ScrewedDifficulty.ToDifficultyDto(),
		FeltLikeDifficulty: *boulder.FeltLikeDifficulty.ToDifficultyDto(),
		Attempts:           boulder.Attempts,
		SessionsTried:      boulder.SessionsTried,
		Exhausting:         boulder.Exhausting,
		Like:               boulder.Like,
		SessionId:          boulder.SessionId,
		Style:              styles,
	}
}
