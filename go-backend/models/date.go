package models

type Date struct {
	day   int
	month int
	year  int
}

func NewDate(day int, month int, year int) *Date {
	return &Date{day: day, month: month, year: year}
}
