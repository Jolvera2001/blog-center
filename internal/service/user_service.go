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
	existingUser, err := s.UserRepo.FindByID(uuid)
	if err != nil {
		return err
	}

	if dto.Name != "" {
		existingUser.Name = dto.Name
	}
	if dto.Email != "" {
		existingUser.Email = dto.Email
	}
	if dto.Password != "" {
		existingUser.Password = dto.Password
	}

	return s.UserRepo.Update(existingUser)
}

func (s *UserService) DeleteUserAccount(uuid string) error {
	return s.UserRepo.Delete(uuid)
}
