package styleservice

import styleRepository "github.com/gorgoroth31/boulder-tracker/go-backend/repository/StyleRepository"

func Add(style string) {
	styleRepository.Add(style)
}
