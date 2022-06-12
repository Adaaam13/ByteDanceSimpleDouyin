package jwtService

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	Username string `json:"username"`
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")
const TokenExpireDuration = time.Hour * 24 * 14

func GenToken(username string, userId uint) (string, error) {
	uc := UserClaims {
		Username: username,
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(TokenExpireDuration),
			Issuer: "simple-tiktok",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, uc)
	s, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return s, nil
}

func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}