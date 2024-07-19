package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
