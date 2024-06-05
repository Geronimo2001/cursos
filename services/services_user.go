package services

import (
	"myapp/dtos"
	"myapp/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) CreateUser(dto dtos.UsersDTO) (*models.User, error) {
	user := &models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Role:     dto.Role,
		Password: dto.Password,
	}
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// putCourse updates a User by ID
func (s *UserService) UpdateUser(id uint, dto dtos.UsersDTO) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	user.Name = dto.Name
	user.Email = dto.Email
	user.Role = dto.Role
	user.Password = dto.Password

	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser deletes a user from the database by ID
func (s *UserService) DeleteUser(id uint) error {
	if err := s.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
