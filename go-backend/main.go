package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/go-backend/models"
)

func main() {
	date := models.NewDate(22, 7, 2024)
	time, _ := models.NewTime(16, 15)
	time2, _ := models.NewTime(18, 15)
	style, _ := models.GetStyleById(3)
	fmt.Print(*style)

	session := models.NewSession(uuid.New(), *date, *time, *time2, false)
	session.Print()
}
