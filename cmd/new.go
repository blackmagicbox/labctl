package cmd

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/blackmagicbox/labctl/internal/tui"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new VM",
	Long:  "Create a new vm from an existing image or from a new one selected by the user",
	Run: func(cmd *cobra.Command, args []string) {
		m := tui.New()
		p := tea.NewProgram(m)

		result, err := p.Run()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error %v\n", err)
		}
		finalModel := result.(tui.Model)
		fmt.Printf("\nVM: \n%v\n", strings.Join(finalModel.Value(), "\n"))

	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
