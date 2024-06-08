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
