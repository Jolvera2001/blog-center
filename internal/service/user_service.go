package service

import "blog-center/internal/domain"

type IUserService interface {
	RegisterUser(name, email, password string) (*domain.User, error)
	GetUserProfile(id uint) (*domain.User, error)
	UpdateUserProfile(name, email, password string) error
	DeleteUserAccount(id uint) error
}