package models

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Name     string
	Status   string `gorm:"default:to do"`
	Reporter string
	Tag      string
	GuildID  string
	UserID   string
}
