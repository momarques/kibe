package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kb/kubecontext"
)

func main() {
	m := kubecontext.NewContextModel()
	program := tea.NewProgram(m)
	if _, err := program.Run(); err != nil {
		panic(err)
	}
}
