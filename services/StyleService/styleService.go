package styleservice

import "github.com/gorgoroth31/boulder-tracker/repository/StyleRepository"

func Add(style string) error {
	err := StyleRepository.Add(style)

	return err
}
