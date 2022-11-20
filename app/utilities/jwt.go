package utilities

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

func SignJWT(claims jwt.Claims, secret string) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(secret))

	return token, err
}

func ParseJWT(claims jwt.Claims, tokenString string, secret string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Forbidden")
		}

		return []byte(secret), nil
	})
}
