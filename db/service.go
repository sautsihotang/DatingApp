package db

import (
	"github.com/sautsihotang/DatingApp/config"
	"github.com/sautsihotang/DatingApp/models"
	"gorm.io/gorm"
)

func SaveUser(user *models.User) error {
	err := config.DB.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
