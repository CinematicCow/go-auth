package auth

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

func GenAuth(email string) string {
	secretKey := []byte("somekeylmao")

	claims := jwt.StandardClaims{
		IssuedAt:  jwt.Now(),
		ExpiresAt: jwt.At(time.Now().Add(time.Minute * 1)),
		Subject:   email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)

	if err != nil {
		log.Fatal("Error at signing token: ", err)
	}

	return signedToken
}
