package repository

import (
	"blog-center/internal/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByID(uuid string) (*domain.User, error) {
	var user domain.User
	if err := r.DB.Where("id = ?", uuid).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *domain.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) Delete(uuid string) error {
	return r.DB.Delete(&domain.User{}, uuid).Error
}
