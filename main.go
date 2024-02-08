package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kb/kubecontext"
	"github.com/momarques/kb/logging"
)

func main() {
	m, err := kubecontext.NewContextModel()
	if err != nil {
		fmt.Printf("failed to create model: %s", err)
		os.Exit(1)
	}
	program := tea.NewProgram(m)

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
