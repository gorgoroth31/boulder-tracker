package models

import "time"

type DateRange struct {
	From time.Time
	To   time.Time
}

type DateRangeDto struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

func NewDateRange(from time.Time, to time.Time) *DateRange {
	return &DateRange{
		From: from,
		To:   to,
	}
}
