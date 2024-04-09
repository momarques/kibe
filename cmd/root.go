package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	core "github.com/momarques/kibe/internal/ui"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "kibe",
	Short: "Kubernetes Interaction with Beauty and Elegancy.",
	Long: `
Kibe aims to be an easy and beautiful tool for interacting with Kubernetes objects on modern terminals.
`,
}

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r", "ru"},
	Short:   "Initialize kibe main UI.",
	Run: func(cmd *cobra.Command, args []string) {
		program := tea.NewProgram(
			core.NewUI(),
			tea.WithAltScreen())

		if _, err := program.Run(); err != nil {
			fmt.Println("failed to run program:", err)
			os.Exit(1)
		}
	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Used for testing layouts without needing to execute the whole program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(testCmd)
}
