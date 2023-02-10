package models

import (

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	EmailID  string
	Name     string
	Password string
	Age      int
}


