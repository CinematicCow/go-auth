package auth

import (
	"log"

	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/xerrors"
)

func ValidateJWT(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("somekeylmao"), nil
	})

	var uErr jwt.UnverfiableTokenError
	var expErr jwt.TokenExpiredError

	if token.Valid {
		return true
	} else if xerrors.As(err, uErr) {
		log.Fatal("This is not even a jwt token bruv.")
		return false
	} else if xerrors.As(err, expErr) {
		log.Fatal("Timing is everything, you are nothing")
		return false
	}
	log.Fatal("Couldn't handle this token")
	return false

}
