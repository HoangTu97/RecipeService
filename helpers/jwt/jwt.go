package jwt

import (
	"Food/domain"
	"Food/helpers/util"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

// GenerateToken generate tokens used for auth
func GenerateToken(userID string, username string, roles []string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := domain.Token{
		UserID: userID,
		Name:   util.EncodeMD5(username),
		Roles:  roles,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*domain.Token, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &domain.Token{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*domain.Token); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
