package service

import "blog-center/internal/domain"

type IUserService interface {
	RegisterUser(name, email, password string) (*domain.User, error)
	GetUserProfile(uuid string) (*domain.User, error)
	UpdateUserProfile(name, email, password string) error
	DeleteUserAccount(uuid string) error
}