package models

import (
	"strconv"
)

type Style struct {
	Id    int
	Alias string
}

func NewStyle(id int, alias string) *Style {
	return &Style{Id: id, Alias: alias}
}

func NewStyleWithoutAlias(id int) *Style {
	return &Style{Id: id, Alias: strconv.Itoa(id)}
}
