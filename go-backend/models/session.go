package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id            uuid.UUID
	Date          Date
	StartTime     Time
	EndTime       Time
	BoulderedSolo bool
	RoutesSolved  []Boulder
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewSession(id uuid.UUID, date Date, startTime Time, endTime Time, boulderedSolo bool, routesSolved []Boulder) *Session {
	return &Session{Date: date, StartTime: startTime, EndTime: endTime, BoulderedSolo: boulderedSolo, RoutesSolved: routesSolved}
}

func (session *Session) Print() {
	fmt.Println(session.Id)
	fmt.Println(session.Date)
	fmt.Println(session.StartTime)
	fmt.Println(session.EndTime)
	fmt.Println(session.DurationInMin())
	fmt.Println(session.BoulderedSolo)
	fmt.Println(session.RoutesSolved)

}

func (session *Session) DurationInMin() int {
	absoluteEndTime := ((session.EndTime.hour + 1) * 60) + session.EndTime.minute
	absoluteStartTime := ((session.StartTime.hour + 1) * 60) + session.StartTime.minute

	return absoluteEndTime - absoluteStartTime
}
