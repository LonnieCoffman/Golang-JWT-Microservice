package daos

import (
	"test/config"
	"test/models"
)

// UserDAO persists user data in database
type UserDAO struct{}

// NewUserDAO creates a new UserDAO
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// Get does the actual query to database, if user with specified id is not found error is returned
func (dao *UserDAO) Get(id uint) (*models.User, error) {
	var user models.User

	err := config.Config.DB.Where("id = ?", id).First(&user).Error

	return &user, err
}
