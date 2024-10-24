package main

import (
	"github.com/gorgoroth31/boulder-tracker/go-backend/repository"
)

func main() {

	repository.OpenConnection()

	// date := models.NewDate(22, 7, 2024)
	// time, _ := models.NewTime(16, 15)
	// time2, _ := models.NewTime(18, 15)

	// difficulties := []models.Difficulty{*models.NewDifficulty(0, "yellow"), *models.NewDifficulty(1, "green"), *models.NewDifficulty(2, "orange"), *models.NewDifficulty(3, "white"), *models.NewDifficulty(4, "blue"), *models.NewDifficulty(5, "red"), *models.NewDifficulty(6, "black")}

	// route1 := models.NewBoudler(difficulties[2], difficulties[2], 1, 1, false, []style.Style{style.Slab, style.Statisch}, true)

	// routes := []models.Boulder{*route1}

	// session := models.NewSession(uuid.New(), *date, *time, *time2, false, routes)

	// fmt.Println(*sessions)
}
