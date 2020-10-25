package models

import (
	"gorm.io/gorm"
)

// UserFeatures entity
type UserFeatures struct {
	gorm.Model
	User       User
}
