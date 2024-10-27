package dto

import "time"

type DateRangeDto struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}
