package userservice

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/models"
	"github.com/gorgoroth31/boulder-tracker/boulder-tracker.api/repository/userrepository"
)

func AddUser(user *models.UserDto) error {
	userEntity := user.ToUserEntity()
	err := userrepository.Add(userEntity)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ExistsUserWithPrincipal(principal string) (bool, error) {
	return userrepository.ExistsUserWithPrincipal(principal)
}

func Delete(userId uuid.UUID) error {
	err := userrepository.Delete(userId)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetByPrincipal(principal string) (*models.UserDto, error) {
	user, err := userrepository.GetByPrincipal(principal)

	if err != nil {
		return nil, err
	}

	userDto := user.ToUserDto()

	return userDto, nil
}
