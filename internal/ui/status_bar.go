package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mistakenelf/teacup/statusbar"
	"github.com/momarques/kibe/internal/ui/style"
)

func newStatusBarModel() statusbar.Model {
	s := statusbar.New(
		style.StatusBarFirstColumnColor(),
		style.StatusBarSecondColumnColor(),
		style.StatusBarThirdColumnColor(),
		style.StatusBarFourthColumnColor(),
	)
	s.SetContent("Resource", "", "", "")
	return s
}

type statusBarUpdated struct{ resource, context, namespace string }

func (m CoreUI) applyStatusBarChanges(msg statusBarUpdated) (CoreUI, tea.Cmd) {
	var cmd tea.Cmd

	m.statusBar.SetContent(
		"Resource", msg.resource,
		fmt.Sprintf("Context: %s", msg.context),
		fmt.Sprintf("Namespace: %s", msg.namespace))

	m.statusBar, cmd = m.statusBar.Update(msg)
	return m, cmd
}

func updateStatusBar(r, c, n string) func() tea.Msg {
	return func() tea.Msg {
		return statusBarUpdated{r, c, n}
	}
}
