package services

import (
	"comp4913-backend/internal/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}

func (s UserService) CreateUser(address string) {
	user := models.User{
		UserAddress: address,
	}
	s.db.FirstOrCreate(&user)
}

func (s UserService) GetUser(address string) *models.User {
	user := models.User{
		UserAddress: address,
	}
	s.db.First(&user)
	return &user
}

func (s UserService) GetUserWithBooks(address string) *models.User {
	user := models.User{
		UserAddress: address,
	}
	s.db.Preload("BoughtBooks").First(&user)
	return &user
}

func (s UserService) GetUserWithPublishes(address string) *models.User {
	user := models.User{
		UserAddress: address,
	}
	s.db.Preload("PublishedBooks").First(&user)
	return &user
}
