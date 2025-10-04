package models

import (
	"time"

	"github.com/google/uuid"
	sessionState "github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/enums"
)

type Session struct {
	Id            uuid.UUID
	StartTime     time.Time
	EndTime       time.Time
	BoulderedSolo bool
	RoutesSolved  []Boulder
	SessionState  sessionState.SessionState
	IsDeleted     bool
	UserId        uuid.UUID
}

type SessionDto struct {
	Id            uuid.UUID                 `json:"id"`
	VisitTime     DateRangeDto              `json:"visitTime"`
	BoulderedSolo bool                      `json:"boulderedSolo"`
	RoutesSolved  []BoulderDto              `json:"routesSolved"`
	SessionState  sessionState.SessionState `json:"sessionState"`
	IsDeleted     bool                      `json:"isDeleted"`
	UserId        uuid.UUID                 `json:"UserId"`
}

func (session *Session) ToSessionDTO() *SessionDto {
	boulderRoutes := []BoulderDto{}
	for _, boulder := range session.RoutesSolved {
		boulderRoutes = append(boulderRoutes, *boulder.ToBoulderDTO())
	}

	visitTime := &DateRangeDto{
		From: session.StartTime,
		To:   session.EndTime,
	}

	return &SessionDto{
		Id:            session.Id,
		VisitTime:     *visitTime,
		BoulderedSolo: session.BoulderedSolo,
		RoutesSolved:  boulderRoutes,
		SessionState:  session.SessionState,
	}
}

func (session *SessionDto) ToSessionEntity() *Session {
	boulderRoutes := []Boulder{}
	for _, boulder := range session.RoutesSolved {
		boulderRoutes = append(boulderRoutes, *boulder.ToBoulderEntity())
	}

	return &Session{
		Id:            session.Id,
		StartTime:     session.VisitTime.From,
		EndTime:       session.VisitTime.To,
		BoulderedSolo: session.BoulderedSolo,
		RoutesSolved:  boulderRoutes,
		SessionState:  session.SessionState,
	}
}
