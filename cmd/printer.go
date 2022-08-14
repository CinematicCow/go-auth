package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Printer = &cobra.Command{
	Use:              "printer",
	Short:            "Prints given text on the screen",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + NAME)
	},
}
