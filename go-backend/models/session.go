package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id            uuid.UUID
	Date          time.Time
	StartTime     time.Time
	EndTime       time.Time
	BoulderedSolo bool
	RoutesSolved  []Boulder
}

func NewSession(id uuid.UUID, date time.Time, startTime time.Time, endTime time.Time, boulderedSolo bool, routesSolved []Boulder) *Session {
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
	time := session.EndTime.Sub(session.StartTime)
	return int(time.Minutes())
}
