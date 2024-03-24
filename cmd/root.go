package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/logging"
	core "github.com/momarques/kibe/internal/ui"
	"github.com/momarques/kibe/internal/ui/style"
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
		// themeConfigFilePath, _ := xdg.ConfigFile("kibe/theme.yaml")
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// fmt.Println(themeConfigFilePath)

		fmt.Println(style.ThemeConfig)

	},
}

func init() {
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(testCmd)
}
