package cmd

import (
	"crypto/rand"
	"fmt"

	"github.com/spf13/cobra"
)

var RandomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random number",
	Long:  "It generates a random number and gives it back. You're free to do anything you want.",
	Run: func(cmd *cobra.Command, args []string) {
		RandomCrypto, _ := rand.Prime(rand.Reader, 10)
		fmt.Println(RandomCrypto)
	},
}
