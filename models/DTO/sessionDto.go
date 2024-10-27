package dto

import (
	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/models"
)

type SessionDto struct {
	Id            uuid.UUID    `json:"id"`
	VisitTime     DateRangeDto `json:"visitTime"`
	BoulderedSolo bool         `json:"boulderedSolo"`
	RoutesSolved  []BoulderDto `json:"routesSolved"`
}

func (session *SessionDto) ToSessionEntity() *models.Session {
	daterange := models.NewDateRange(session.VisitTime.From, session.VisitTime.To)
	boulderRoutes := []models.Boulder{}
	for _, boulder := range session.RoutesSolved {
		boulderRoutes = append(boulderRoutes, *boulder.ToBoulderEntity())
	}

	return &models.Session{
		Id:            session.Id,
		VisitTime:     *daterange,
		BoulderedSolo: session.BoulderedSolo,
		RoutesSolved:  boulderRoutes,
	}
}
