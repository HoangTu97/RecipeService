package models

import (
	"gorm.io/gorm"
	"time"
)

// Location entity
type Location struct {
	gorm.Model
	Name      string  `gorm:"type:varchar(255"`
	Address   string  `gorm:"type:varchar(255)"`
	Long      float32 `gorm:"index:location"`
	Lat       float32 `gorm:"index:location"`
	OpenHour  time.Time
	CloseHour time.Time
}
