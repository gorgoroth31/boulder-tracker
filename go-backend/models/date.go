package models

import "github.com/google/uuid"

type Date struct {
	Id        uuid.UUID
	day       int
	month     int
	year      int
	SessionId uint
}

func NewDate(day int, month int, year int) *Date {
	return &Date{day: day, month: month, year: year}
}
