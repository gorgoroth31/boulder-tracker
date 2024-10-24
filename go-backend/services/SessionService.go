package services

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
