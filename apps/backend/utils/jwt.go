package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
}

var ErrNoSetSecretKey = errors.New("JWT_SECRET_KEY is not set")
var ErrFailedSignToken = errors.New("failed to sign token")

func NewToken(userID int64) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatInt(userID, 10),
			Issuer:    "adpay",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", ErrNoSetSecretKey
	}
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", ErrFailedSignToken
	}

	return tokenString, nil

}
