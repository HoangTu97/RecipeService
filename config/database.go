package config

import (
  "p2/models"
)

func GetModelsNeedMigrate() []interface{} {
  return []interface{}{
    // Models declare
    &models.User{},
    &models.Link{},
    &models.Click{},
    // Models declare end : dont remove
  }
}
