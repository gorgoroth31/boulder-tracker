package styleservice

import "github.com/gorgoroth31/boulder-tracker/repository/stylerepository"

func Add(style string) error {
	err := stylerepository.Add(style)

	return err
}
