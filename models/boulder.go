package models

import "github.com/google/uuid"

type Boulder struct {
	Id                 uuid.UUID
	ScrewedDifficulty  Difficulty
	FeltLikeDifficulty Difficulty
	Attempts           int
	SessionsTried      int
	Exhausting         bool
	Style              []Style
	Like               bool
	SessionId          uuid.UUID
}

func NewBoudler(Id uuid.UUID, screwedDifficulty Difficulty,
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
