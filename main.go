package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	kubecontext "github.com/momarques/kibe/kube/context"
	"github.com/momarques/kibe/logging"
	listmodel "github.com/momarques/kibe/model/list"
)

func main() {
	initialModel, err := listmodel.New(kubecontext.NewContextBindings())
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
}
