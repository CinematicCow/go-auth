package cmd

import (
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

var VerifyJwt = &cobra.Command{
	Use:              "verify",
	Short:            "Verifies JWT generated by login",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {

		token, err := jwt.Parse(TOKEN, func(token *jwt.Token) (interface{}, error) {
			return []byte("secretkeylul"), nil
		})

		fmt.Printf("val %v\ndan %d", *token, &token)
		var uErr *jwt.UnverfiableTokenError
		var expErr *jwt.TokenExpiredError
		var nbfErr *jwt.TokenNotValidYetError

		// Use xerrors.Is to see what kind of error(s) occurred
		if token.Valid {
			fmt.Println("You look nice today")
		} else if xerrors.As(err, &uErr) {
			fmt.Println("That's not even a token")
		} else if xerrors.As(err, &expErr) {
			fmt.Println("Timing is everything")
		} else if xerrors.As(err, &nbfErr) {
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	},
}
