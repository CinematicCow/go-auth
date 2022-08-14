package auth

import (
	"github.com/dgrijalva/jwt-go/v4"
)

func ValidateJWT(tokenString string) bool {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("somekeylmao"), nil
	})

	if err != nil {
		return false
	}
	if token.Valid {
		return true
	} else {
		return false
	}

}
