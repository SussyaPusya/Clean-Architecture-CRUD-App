package entity

import "gorm.io/gorm"

type User struct {
	UserName string `json:"name"`
	Password string `json:"pass"`
	Auth     string `json:"auth"`

	BirhDay string `json:"bithday"`
}

type UserDb struct {
	gorm.Model
	UserName string `gorm:"size:255"`
	Password string `gorm:"size:255"`

	BirhDay string `gorm:"size:255"`
}
