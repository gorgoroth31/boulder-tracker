package sessioncontroller

import (
	"fmt"

	"github.com/gorgoroth31/boulder-tracker/models"
)

func AddSession(session *models.Session) {
	fmt.Print(session)
}
