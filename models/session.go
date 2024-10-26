package models

import (
	"github.com/google/uuid"
)

type Session struct {
	Id            uuid.UUID
	VisitTime     DateRange
	BoulderedSolo bool
	RoutesSolved  []Boulder
}
