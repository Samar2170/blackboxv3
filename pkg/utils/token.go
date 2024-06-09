package utils

import (
	"blackboxv3/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var signingKey []byte

func init() {
	signingKey = []byte(config.Config.SigningKey)
}

type UserClaim struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateToken(username string, userID string) (string, error) {
	claims := UserClaim{
		username,
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "cognitio",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(tokenString string) (*UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserClaim)
	if !ok {
		return nil, err
	}
	if err := claims.Valid(); err != nil {
		return nil, err
	}
	if claims.RegisteredClaims.ExpiresAt.Time.Before(time.Now()) {
		return nil, jwt.ErrTokenExpired
	}
	return claims, nil
}
