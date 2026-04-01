package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "labctl",
	Short: "Create and manage VMs",
	Long:  "Create and manage your vms leveraging the cloud-init interface",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
