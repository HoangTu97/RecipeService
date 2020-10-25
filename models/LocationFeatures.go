package models

import (
	"gorm.io/gorm"
)

// LocationFeatures entity
type LocationFeatures struct {
	gorm.Model
	Location Location
}
