package services

import "github.com/gorgoroth31/boulder-tracker/models"

func SampleSession() *models.Session {
	// date := models.NewDate(22, 7, 2024)
	// time, _ := models.NewTime(16, 15)
	// time2, _ := models.NewTime(18, 15)

	// difficulties := []models.Difficulty{*models.NewDifficulty(0, "yellow"), *models.NewDifficulty(1, "green"), *models.NewDifficulty(2, "orange"), *models.NewDifficulty(3, "white"), *models.NewDifficulty(4, "blue"), *models.NewDifficulty(5, "red"), *models.NewDifficulty(6, "black")}

	//route1 := models.NewBoudler(uuid.New(), difficulties[2], difficulties[2], 1, 1, false, []style.Style{style.Slab, style.Statisch}, true)

	// routes := []models.Boulder{*route1}

	// session := models.NewSession(uuid.New(), *date, *time, *time2, false, routes)
	return &models.Session{}
}

// func AddSession(session *models.Session) error {
// 	db := repository.DB

// 	// hier müssen die anderen Proeprties manuell in die DB gesetzt werden

// 	result := db.Create(&session)

// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

// func DeleteSession(id uuid.UUID) {
// 	fmt.Println("about to delete id ", id)
// 	// cascade delete boulder entities
// }

// func GetSessions() *[]models.Session {
// 	fmt.Println("returning sessions")
// 	sessions := &[]models.Session{}
// 	result := repository.DB.Find(&sessions)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return sessions
// }
