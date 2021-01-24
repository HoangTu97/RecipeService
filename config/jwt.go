package config

import (
  "Food/helpers/jwt"
  "Food/helpers/setting"
)

func SetupJWT(appSetting setting.App) jwt.JwtManager {
  return jwt.NewJwtManager(appSetting)
}