package service

import (
	"blog-center/internal/domain"

	"google.golang.org/grpc/resolver/passthrough"
)

type UserService struct {
	UserRepo domain.IUserRepository
}

func (s *UserService) RegisterUser(name, email, password string) (*domain.User, error) {

}

func (s *UserService) GetUserProfile(id uint) (*domain.User, error) {

}

func (s *UserService) UpdateUserProfile(name, email, password string)error {

}

func (s *UserService) DeleteUserAccount(id uint) error {

}



