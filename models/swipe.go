package models

import (
	"time"

	"gorm.io/gorm"
)

type Swipe struct {
	gorm.Model
	SwipeID   int       `json:"swipeId"`
	UserID    int       `json:"userId"`
	ProfileID int       `json:"profilId"`
	SwipeType string    `json:"swipeType"`
	Timestamp time.Time `json:"timetamp"`

	User    User
	Profile Profile
}
