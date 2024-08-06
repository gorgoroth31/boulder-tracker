package models

import "errors"

type Time struct {
	hour   int
	minute int
}

func NewTime(hour int, minute int) (*Time, error) {
	if hour >= 24 || hour < 0 || minute >= 60 || minute < 0 {
		return nil, errors.New("invalid time format")
	}
	return &Time{hour: hour, minute: minute}, nil
}
