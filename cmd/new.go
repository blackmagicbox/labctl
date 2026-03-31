package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new VM",
	Long:  "Create a new vm from an existing image or from a new one selected by the user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating a new vm...")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
