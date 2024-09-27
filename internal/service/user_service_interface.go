package service

import (
	"blog-center/internal/domain"
	"blog-center/internal/dtos"
)

type IUserService interface {
	RegisterUser(dto dtos.UserDto) (*domain.User, error)
	GetUserProfile(uuid string) (*domain.User, error)
	UpdateUserProfile(uuid string, dto dtos.UserDto) error
	DeleteUserAccount(uuid string) error
}