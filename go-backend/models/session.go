package models

import (
	"fmt"

	"github.com/google/uuid"
)

type Session struct {
	id            uuid.UUID
	date          Date
	startTime     Time
	endTime       Time
	boulderedSolo bool
	routesSolved  []Boulder
}

func NewSession(id uuid.UUID, date Date, startTime Time, endTime Time, boulderedSolo bool, routesSolved []Boulder) *Session {
	return &Session{id: id, date: date, startTime: startTime, endTime: endTime, boulderedSolo: boulderedSolo, routesSolved: routesSolved}
}

func (session *Session) Print() {
	fmt.Println(session.id)
	fmt.Println(session.date)
	fmt.Println(session.startTime)
	fmt.Println(session.endTime)
	fmt.Println(session.DurationInMin())
	fmt.Println(session.boulderedSolo)
	fmt.Println(session.routesSolved)

}

func (session *Session) DurationInMin() int {
	absoluteEndTime := ((session.endTime.hour + 1) * 60) + session.endTime.minute
	absoluteStartTime := ((session.startTime.hour + 1) * 60) + session.startTime.minute

	return absoluteEndTime - absoluteStartTime
}
