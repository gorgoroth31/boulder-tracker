package models

import (
	"github.com/google/uuid"
)

type Style struct {
	Id    uuid.UUID
	Alias string
}

type StyleDto struct {
	Id    uuid.UUID `json:"id"`
	Alias string    `json:"alias"`
}

func NewStyle(id uuid.UUID, alias string) *Style {
	return &Style{Id: id, Alias: alias}
}

func (dto *StyleDto) ToStyleEntity() *Style {
	return &Style{
		Id:    dto.Id,
		Alias: dto.Alias,
	}
}

func (entity *Style) ToStyleDTO() *StyleDto {
	return &StyleDto{
		Id:    entity.Id,
		Alias: entity.Alias,
	}
}
