package cmd

import (
	"fmt"
	"os"

	"github.com/blackmagicbox/labctl/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new VM",
	Long:  "Create a new vm from an existing image or from a new one selected by the user",
	Run: func(cmd *cobra.Command, args []string) {
		m := tui.New("rocky-9-20260331")
		p := tea.NewProgram(m)

		result, err := p.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error %v\n", err)
		}
		finalModel := result.(tui.Model)
		fmt.Printf("VM name: %s\n", finalModel.Value())

	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
