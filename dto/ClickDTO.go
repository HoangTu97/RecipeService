package dto

import (
  "time"
  "gorm.io/gorm"
)

// ClickDTO godoc
type ClickDTO struct {
  ID uint `json:"id,omitempty"`
  CreatedAt time.Time `json:"created_at,omitempty"`
  UpdatedAt time.Time `json:"updated_at,omitempty"`
  DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

  Ip string `json:"ip,omitempty"`
  Country string `json:"country,omitempty"`
  Referer string `json:"referer,omitempty"`
  RefererHost string `json:"referer_host,omitempty"`
  UserAgent string `json:"user_agent,omitempty"`

  // Auto-gen relationships
  LinkID uint `json:"link_id,omitempty"`
  // Auto-gen relationships : dont remove
}
