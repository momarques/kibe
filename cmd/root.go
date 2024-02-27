package cmd

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/logging"
	"github.com/momarques/kibe/internal/ui"
	core "github.com/momarques/kibe/internal/ui"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "kibe",
	Short: "Kubernetes Interaction with Beauty and Elegancy.",
	Long: `
Kibe aims to be a easy tool for interacting with Kubernetes objects without showing a lot of hard to understand information. 
It's a tool focused for developers who doesn't necessarily need to understand the internals of Kubernetes resources.
Also it's a tool made to look beautiful on modern terminals.
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

		if len(os.Getenv("DEBUG")) > 0 {
			f, err := tea.LogToFile(logging.LogFile, "debug")
			if err != nil {
				fmt.Println("failed to set log file:", err)
				os.Exit(1)
			}
			defer f.Close()
		}

		if _, err := program.Run(); err != nil {
			fmt.Println("failed to run program:", err)
			os.Exit(1)
		}
	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Used for testing layouts without needing to execute the whole program",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(ui.NewModel())
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(testCmd)
}
