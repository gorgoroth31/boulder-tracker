package models

import (
	"github.com/google/uuid"
)

type Difficulty struct {
	Id            uuid.UUID
	Alias         string
	RelativeLevel int
}

type DifficultyDto struct {
	Id            uuid.UUID `json:"id"`
	Alias         string    `json:"alias"`
	RelativeLevel int       `json:"relativeLevel"`
}

func NewDifficulty(Id uuid.UUID, Alias string, RelativeLevel int) *Difficulty {
	return &Difficulty{
		Id:            Id,
		Alias:         Alias,
		RelativeLevel: RelativeLevel,
	}
}

func (dto *DifficultyDto) ToDifficultyEntity() *Difficulty {
	return &Difficulty{
		Id:            dto.Id,
		Alias:         dto.Alias,
		RelativeLevel: dto.RelativeLevel,
	}
}

func (entity *Difficulty) ToDifficultyDto() *DifficultyDto {
	return &DifficultyDto{
		Id:            entity.Id,
		Alias:         entity.Alias,
		RelativeLevel: entity.RelativeLevel,
	}
}
