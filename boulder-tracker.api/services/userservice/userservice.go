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

func Delete(userId uuid.UUID) error {
	err := userrepository.Delete(userId)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetByEmail(userEmail string) (*models.UserDto, error) {
	user, err := userrepository.GetByEmail(userEmail)

	if err != nil {
		return nil, err
	}

	userDto := user.ToUserDto()

	return userDto, nil
}
