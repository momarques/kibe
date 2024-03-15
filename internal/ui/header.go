package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

type headerTitleUpdated string
type headerItemCountUpdated int

type headerModel struct {
	text      headerTitleUpdated
	itemCount headerItemCountUpdated
}

func (c *tableContent) updateHeader(itemCount int) tea.Cmd {
	return func() tea.Msg {
		return headerItemCountUpdated(itemCount)
	}
}

func (s *listSelector) updateHeader(title string) tea.Cmd {
	return func() tea.Msg {
		return headerTitleUpdated(title)
	}
}

func (m CoreUI) headerView() string {
	titleStyle := uistyles.
		ViewTitleStyle.
		Copy().
		PaddingLeft(1).
		MarginTop(1).
		MarginLeft(2)

	itemCountStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#685e59")).
		MarginTop(1).
		MarginLeft(2).
		MarginBottom(1)

	return lipgloss.JoinVertical(lipgloss.Left,
		titleStyle.Render(string(m.header.text)),
		itemCountStyle.Render(fmt.Sprintf("%d items", m.header.itemCount)),
	)
}
