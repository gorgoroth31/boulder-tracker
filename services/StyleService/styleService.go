package styleservice

import stylerepository "github.com/gorgoroth31/boulder-tracker/repository/styleRepository"

func Add(style string) error {
	err := stylerepository.Add(style)

	return err
}
