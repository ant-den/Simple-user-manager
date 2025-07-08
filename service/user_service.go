package service

import (
	"GO_Project/models"
	"GO_Project/repository"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}
func (s *userService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}
func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}
func (s *userService) GetUserByID(id uint) (models.User, error) {
	return s.repo.FindByID(id)
}
func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}
func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
