package models

import (
	dr "github.com/felixenescu/date-range"
	"github.com/google/uuid"
)

type Session struct {
	Id            uuid.UUID
	VisitTime     dr.DateRange
	BoulderedSolo bool
	RoutesSolved  []Boulder
}
