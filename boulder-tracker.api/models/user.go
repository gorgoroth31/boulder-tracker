package models

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID
	UserName  string
	Email     string
	IsDeleted bool
	Sessions  []Session
}

type UserDto struct {
	Id        uuid.UUID `json:"id"`
	UserName  string    `json:"userName"`
	Email     string    `json:"email"`
	IsDeleted bool      `json:"isDeleted"`
}

func (user *User) ToUserDto() *UserDto {
	return &UserDto{
		Id:        user.Id,
		UserName:  user.UserName,
		Email:     user.Email,
		IsDeleted: user.IsDeleted,
	}
}

func (user *UserDto) ToUserEntity() *User {
	return &User{
		Id:        user.Id,
		UserName:  user.UserName,
		Email:     user.Email,
		IsDeleted: user.IsDeleted,
	}
}
