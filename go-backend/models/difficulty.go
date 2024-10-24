package models

import (
	"strconv"
)

type Difficulty struct {
	Id    int
	alias string
}

func NewDifficulty(id int, alias string) *Difficulty {
	return &Difficulty{Id: id, alias: alias}
}

func NewDifficultyWithoutAlias(id int) *Difficulty {
	return &Difficulty{Id: id, alias: strconv.Itoa(id)}
}
