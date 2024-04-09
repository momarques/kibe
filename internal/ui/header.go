package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/ui/style"
)

type headerTitleUpdated string
type headerItemCountUpdated int

type headerModel struct {
	text      headerTitleUpdated
	itemCount headerItemCountUpdated
}

func updateHeaderItemCount(itemCount int) tea.Cmd {
	return func() tea.Msg {
		return headerItemCountUpdated(itemCount)
	}
}

func updateHeaderTitle(title string) tea.Cmd {
	return func() tea.Msg {
		return headerTitleUpdated(title)
	}
}

func (m CoreUI) headerView() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		style.CoreHeaderTitleStyle().Render(string(m.header.text)),
		style.CoreHeaderItemCountStyle().Render(fmt.Sprintf("%d items", m.header.itemCount)),
	)
}
