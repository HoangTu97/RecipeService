package models

import (
  "gorm.io/gorm"
)

// Click entity
type Click struct {
  gorm.Model

  // Auto-gen fields
  Ip string `gorm:"type:varchar(255)"`
  Country string `gorm:"type:varchar(255)"`
  Referer string `gorm:"type:varchar(255)"`
  RefererHost string `gorm:"type:varchar(255)"`
  UserAgent string `gorm:"type:text"`
  // Auto-gen fields : dont remove

  // Auto-gen relationships
  LinkID uint
  Link Link
  // Auto-gen relationships : dont remove
}
