package dto

import (
  "time"
  "gorm.io/gorm"
)

// LinkDTO godoc
type LinkDTO struct {
  ID uint `json:"id,omitempty"`
  CreatedAt time.Time `json:"created_at,omitempty"`
  UpdatedAt time.Time `json:"updated_at,omitempty"`
  DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

  ShortUrl string `json:"short_url,omitempty"`
  LongUrl string `json:"long_url,omitempty"`
  LongUrlHash string `json:"long_url_hash,omitempty"`
  Ip string `json:"ip,omitempty"`
  ClickNum uint `json:"click_num,omitempty"`
  SecretKey string `json:"secret_key,omitempty"`
  IsDisabled bool `json:"is_disabled,omitempty"`
  IsCustom bool `json:"is_custom,omitempty"`
  IsApi bool `json:"is_api,omitempty"`

  // Auto-gen relationships
  CreatorID uint `json:"creator_id,omitempty"`
  // Auto-gen relationships : dont remove
}
