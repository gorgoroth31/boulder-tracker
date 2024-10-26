package models

import "time"

type DateRange struct {
	From time.Time
	To   time.Time
}

func NewDateRange(from time.Time, to time.Time) *DateRange {
	return &DateRange{
		From: from,
		To:   to,
	}
}
