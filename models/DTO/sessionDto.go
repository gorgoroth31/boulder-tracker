package dto

import (
	"time"

	"github.com/google/uuid"
)

type SessionDto struct {
	Id            uuid.UUID    `json:"id"`
	Date          time.Time    `json:"date"`
	StartTime     time.Time    `json:"startTime"`
	EndTime       time.Time    `json:"endTime"`
	BoulderedSolo bool         `json:"boulderedSolo"`
	RoutesSolved  []BoulderDto `json:"routesSolved"`
}
