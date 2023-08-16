package services

import (
	"smartPost/models"
	"smartPost/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepository.Create(user)
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.UserRepository.GetUsers()
}

func (s *UserService) UpdateName(id int, newUserName string) (models.User, error) {
	return s.UserRepository.UpdateName(id, newUserName)
}
func (s *UserService) UpdateEmail(id int, newUserEmail string) (models.User, error) {
	return s.UserRepository.UpdateEmail(id, newUserEmail)
}
func (s *UserService) DeleteUser(id int) error {
	return s.UserRepository.DeleteUser(id)
}
func (s *UserService) IsEmailTaken(email string) bool {
	return s.UserRepository.IsEmailTaken(email)
}
