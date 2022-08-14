package cmd

import (
	"encoding/json"
	"fmt"
	"gauth/auth"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func checkUserExists(username *string, password *string) (User, bool) {
	content, err := os.ReadFile("db/db.json")
	// byteValue, err := os.ReadFile("db/db.json")
	if err != nil {
		log.Fatal("Error when reading database: ", err)
	}

	var user User

	err = json.Unmarshal(content, &user)
	if err != nil {
		log.Fatal("Error at Unmarshal() ", err)
	}

	// fmt.Printf("DB\nusername: %s | password: %s\nCLI\nusername: %v | password: %v\n", user.Username, user.Password, *username, *password)

	if user.Username == *username && user.Password == *password {
		return user, true
	} else {
		return user, false
	}

}

var Login = &cobra.Command{
	Use:              "login",
	Short:            "Creates JWT key",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		user, isUser := checkUserExists(&USERNAME, &PASSWORD)
		if !isUser {
			log.Fatal("User not found")
		}
		token := auth.GenAuth(user.Username)
		fmt.Println("token: ", token)
	},
}
