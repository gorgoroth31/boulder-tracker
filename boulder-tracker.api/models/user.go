package models

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID
	UserName  string
	Principal string
	IsDeleted bool
	Sessions  []Session
}

type UserDto struct {
	Id        uuid.UUID `json:"id"`
	UserName  string    `json:"userName"`
	Principal string    `json:"principal"`
	IsDeleted bool      `json:"isDeleted"`
}

func (user *User) ToUserDto() *UserDto {
	return &UserDto{
		Id:        user.Id,
		UserName:  user.UserName,
		Principal: user.Principal,
		IsDeleted: user.IsDeleted,
	}
}

func (user *UserDto) ToUserEntity() *User {
	return &User{
		Id:        user.Id,
		UserName:  user.UserName,
		Principal: user.Principal,
		IsDeleted: user.IsDeleted,
	}
}
