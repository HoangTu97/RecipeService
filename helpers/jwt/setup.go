package jwt

import (
	"Food/config"
)

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(config.AppSetting.JwtSecret)
}
