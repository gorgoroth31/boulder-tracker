package models

import (
	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/go-backend/enums/style"
)

type Boulder struct {
	Id                 uuid.UUID
	screwedDifficulty  Difficulty
	feltLikeDifficulty Difficulty
	attempts           int
	sessionsTried      int
	exhausting         bool
	style              []style.Style
	like               bool
	SessionID          uuid.UUID
}

func NewBoudler(Id uuid.UUID, screwedDifficulty Difficulty,
	feltLikeDifficulty Difficulty,
	attempts int,
	sessionsTried int,
	exhausting bool,
	style []style.Style, like bool) *Boulder {
	return &Boulder{
		Id:                 Id,
		screwedDifficulty:  screwedDifficulty,
		feltLikeDifficulty: feltLikeDifficulty,
		attempts:           attempts,
		sessionsTried:      sessionsTried,
		exhausting:         exhausting,
		style:              style,
		like:               like,
	}
}
