package main

import (
	Dbcontext "github.com/gorgoroth31/boulder-tracker/go-backend/repository/DbContext"
	styleservice "github.com/gorgoroth31/boulder-tracker/go-backend/services/StyleService"
)

func main() {

	Dbcontext.OpenConnection()

	styleservice.Add("newStyle")

	// date := models.NewDate(22, 7, 2024)
	// time, _ := models.NewTime(16, 15)
	// time2, _ := models.NewTime(18, 15)

	// difficulties := []models.Difficulty{*models.NewDifficulty(0, "yellow"), *models.NewDifficulty(1, "green"), *models.NewDifficulty(2, "orange"), *models.NewDifficulty(3, "white"), *models.NewDifficulty(4, "blue"), *models.NewDifficulty(5, "red"), *models.NewDifficulty(6, "black")}

	//route1 := models.NewBoudler(uuid.New(), difficulties[2], difficulties[2], 1, 1, false, []style.Style{style.Slab, style.Statisch}, true)

	// routes := []models.Boulder{*route1}

	// session := models.NewSession(uuid.New(), *date, *time, *time2, false, routes)

	// fmt.Println(*sessions)
}
