package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	ProfileID int    `json:"profileId"`
	UserID    int    `json:"userId"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Location  string `json:"location"`
}
