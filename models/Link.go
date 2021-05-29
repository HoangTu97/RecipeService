package models

import (
  "gorm.io/gorm"
)

// Link entity
type Link struct {
  gorm.Model

  // Auto-gen fields
  ShortUrl string `gorm:"type:varchar(255)"`
  LongUrl string `gorm:"type:text"`
  LongUrlHash string `gorm:"type:varchar(255)"`
  Ip string `gorm:"type:varchar(255)"`
  ClickNum uint
  SecretKey string `gorm:"type:varchar(255)"`
  IsDisabled bool
  IsCustom bool
  IsApi bool
  // Auto-gen fields : dont remove

  // Auto-gen relationships
  Clicks []Click
  CreatorID uint
  Creator User
  // Auto-gen relationships : dont remove
}
