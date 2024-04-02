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

func (c tableContent) updateHeader(itemCount int) tea.Cmd {
	return func() tea.Msg {
		return headerItemCountUpdated(itemCount)
	}
}

func (c clientConfigSelector) updateHeader(title string) tea.Cmd {
	return func() tea.Msg {
		return headerTitleUpdated(title)
	}
}

func (m CoreUI) headerView() string {
	titleStyle := style.CoreHeaderTitleStyle
	itemCountStyle := style.CoreHeaderItemCountStyle

	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle().Render(string(m.header.text)),
		itemCountStyle().Render(fmt.Sprintf("%d items", m.header.itemCount)),
	)
}
