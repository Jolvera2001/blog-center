package service

import (
	"blog-center/internal/domain"
	"blog-center/internal/dtos"
)

type UserService struct {
	UserRepo domain.IUserRepository
}

func (s *UserService) RegisterUser(dto dtos.UserDto) (string, error) {
	newUser := domain.User{
		Name: dto.Name,
		Email: dto.Email,
		Password: dto.Password,
	}

	err := s.UserRepo.Create(&newUser)
	if err != nil {
		return "", err
	}

	return newUser.ID, nil
}

func (s *UserService) GetUserProfile(id string) (*domain.User, error) {
	foundUser, err := s.UserRepo.FindByID(id)
	if err != nil {
		return &domain.User{}, err
	}

	return foundUser, nil
}

func (s *UserService) UpdateUserProfile(uuid string, dto dtos.UserDto) error {

}

func (s *UserService) DeleteUserAccount(uuid string) error {

}
