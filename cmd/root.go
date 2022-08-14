package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Auther",
	Short: "Auther is a jwt implementation in guu",
}

var NAME string
var USERNAME string
var PASSWORD string
var TOKEN string

func init() {
	rootCmd.AddCommand(RandomCmd)
	Printer.Flags().StringVarP(&NAME, "name", "n", "", "name to print")
	rootCmd.AddCommand(Printer)
	Login.Flags().StringVarP(&USERNAME, "username", "n", "", "your username")
	Login.Flags().StringVarP(&PASSWORD, "password", "p", "", "your password")
	rootCmd.AddCommand(Login)
	VerifyJwt.Flags().StringVarP(&TOKEN, "token", "t", "", "your JWT token")
	rootCmd.AddCommand(VerifyJwt)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
