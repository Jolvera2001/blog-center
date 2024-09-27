package service

import "blog-center/internal/domain"

type UserService struct {
	UserRepo domain.IUserRepository
}