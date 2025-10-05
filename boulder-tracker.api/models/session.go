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
	StartTime     time.Time                 `json:"startTime"`
	EndTime       time.Time                 `json:"endTime"`
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

	return &SessionDto{
		Id:            session.Id,
		StartTime:     session.StartTime,
		EndTime:       session.EndTime,
		BoulderedSolo: session.BoulderedSolo,
		RoutesSolved:  boulderRoutes,
		IsDeleted:     session.IsDeleted,
		SessionState:  session.SessionState,
		UserId:        session.UserId,
	}
}

func (session *SessionDto) ToSessionEntity() *Session {
	boulderRoutes := []Boulder{}
	for _, boulder := range session.RoutesSolved {
		boulderRoutes = append(boulderRoutes, *boulder.ToBoulderEntity())
	}

	return &Session{
		Id:            session.Id,
		StartTime:     session.StartTime,
		EndTime:       session.EndTime,
		BoulderedSolo: session.BoulderedSolo,
		RoutesSolved:  boulderRoutes,
		SessionState:  session.SessionState,
	}
}
