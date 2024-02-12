package tablemodel

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/kube"
	"github.com/momarques/kibe/internal/logging"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func Table() {
	c := kube.NewKubeClient("kind-kibe")

	m, err := New("Namespace", "kube-system", c)
	if err != nil {
		logging.Log.Error(err)
	}

	p := tea.NewProgram(m)
	_, err = p.Run()
	if err != nil {
		logging.Log.Error(err)
	}
}
