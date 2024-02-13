package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

	// func Table() {
	// 	c := kube.NewKubeClient("kind-kibe")

	// 	m, err := New("Service", "kube-system", c)
	// 	if err != nil {
	// 		logging.Log.Error(err)
	// 	}

	// 	p := tea.NewProgram(m)
	// 	_, err = p.Run()
	// 	if err != nil {
	// 		logging.Log.Error(err)
	// 	}
	// }
