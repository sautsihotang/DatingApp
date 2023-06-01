package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID        int    `json:"userId"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	PremiumStatus bool   `json:"premiumStatus" gorm:"default:false"`
}
