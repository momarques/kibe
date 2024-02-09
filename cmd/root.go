package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/kube/contextactions"
	kubeapis "github.com/momarques/kibe/internal/kube/resources"
	"github.com/momarques/kibe/internal/logging"
	listmodel "github.com/momarques/kibe/internal/model/list"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "kibe",
	Short: "Kubernetes Interacvtive, Beautiful and Easy CLI for managing Kubernetes resources.",
	Long: `
Kibe aims to be a easy tool for interacting with Kubernetes objects without showing a lot of hard to understand information. 
It's a tool focused for developers who doesn't necessarily need to understand the internals of Kubernetes resources.
Also it's a tool made to look beautiful on modern terminals.
`,
	Run: func(cmd *cobra.Command, args []string) {
		initialModel, err := listmodel.New(contextactions.New())
		if err != nil {
			fmt.Printf("failed to create model: %s", err)
			os.Exit(1)
		}
		program := tea.NewProgram(initialModel)

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
	Short: "Used for calling internal functions to test CLI behavior",
	Long: `
If you need to add a new command or a new process to an existing command, 
the 'test' command may be useful to execute the new flow isolated.
Just put what you need to test inside the Run field and test with:

	make test-command`,
	Run: func(cmd *cobra.Command, args []string) {
		kubeapis.AAAA()
	},
}

func init() {
	RootCmd.AddCommand(testCmd)
}
