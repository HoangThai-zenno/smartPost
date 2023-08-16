package repositories

import (
	"github.com/jinzhu/gorm"
	"smartPost/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := r.DB.Table("users").Order("id").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) UpdateName(id int, newUserName string) (models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	user.Name = newUserName
	if err := r.DB.Save(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (r *UserRepository) UpdateEmail(id int, newUserEmail string) (models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	user.Email = newUserEmail
	if err := r.DB.Save(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (r *UserRepository) DeleteUser(id int) error {
	var user models.User
	if err := r.DB.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) IsEmailTaken(email string) bool {
	var count int64
	r.DB.Table("users").Where("email = ?", email).Count(&count)
	return count > 0
}
