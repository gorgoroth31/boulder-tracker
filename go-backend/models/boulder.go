package models

import (
	"github.com/gorgoroth31/boulder-tracker/go-backend/enums/style"
)

type Boulder struct {
	screwedDifficulty  Difficulty
	feltLikeDifficulty Difficulty
	attempts           int
	sessionsTried      int
	exhausting         bool
	style              []style.Style
	like               bool
}

func NewBoudler(screwedDifficulty Difficulty,
	feltLikeDifficulty Difficulty,
	attempts int,
	sessionsTried int,
	exhausting bool,
	style []style.Style, like bool) *Boulder {
	return &Boulder{
		screwedDifficulty:  screwedDifficulty,
		feltLikeDifficulty: feltLikeDifficulty,
		attempts:           attempts,
		sessionsTried:      1,
		exhausting:         exhausting,
		style:              style,
		like:               like,
	}
}
