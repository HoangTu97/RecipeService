package jwt

import (
	"Food/helpers/setting"
)

// Setup Initialize the util
func Setup(appSetting setting.App) {
	jwtSecret = []byte(appSetting.JwtSecret)
}
