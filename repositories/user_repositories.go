package repositories

import (
	"fmt"
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
	fmt.Println("User Respositories: ", user)
	return r.DB.Create(user).Error
}
func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := r.DB.Preload("Groups").Order("id").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (r *UserRepository) GetUserGroups() ([]models.UserGroup, error) {
	var userGroups []models.UserGroup
	result := r.DB.Find(&userGroups)
	return userGroups, result.Error
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
func (r *UserRepository) UpdateRole(id int, newUserRole string) (models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	user.Role = newUserRole
	if err := r.DB.Save(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (r *UserRepository) CreateUserGroup(groups []models.UserGroup) error {
	//r.DB.Delete()
	//var arr = groups
	//fmt.Println(groups)
	//return r.DB.Create(&groups).Error
	for i := 0; i < len(groups); i++ {
		r.DB.Create(groups[i])
	}
	return nil
	//payload := []*models.UserGroup{
	//	{
	//		UserId:  1,
	//		GroupId: 1,
	//	},
	//	{
	//		UserId:  5,
	//		GroupId: 2,
	//	},
	//}
	//result := r.DB.Create(payload)
	//return result.Error
}
func (r *UserRepository) DeleteUserGroup(userId int) error {
	var userGroup []models.UserGroup
	if err := r.DB.Where("user_id = ?", userId).Delete(&userGroup).Error; err != nil {
		return err
	}

	return nil
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
func (r *UserRepository) GetGroups() ([]models.Group, error) {
	var groups []models.Group
	r.DB.Find(&groups)
	return groups, nil
}
