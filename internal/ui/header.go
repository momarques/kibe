package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

type headerModel struct {
	text      string
	itemCount string
}

type headerUpdated struct {
	text      string
	itemCount string
}

func (c *content) updateHeader(title string, itemCount int) tea.Cmd {
	return func() tea.Msg {
		return headerUpdated{
			text:      title,
			itemCount: fmt.Sprintf("%d items", itemCount)}
	}
}

func (t headerModel) viewHeaderUI(size int) string {
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
		MarginBottom(3)

	return lipgloss.JoinVertical(lipgloss.Left,
		titleStyle.Render(t.text),
		itemCountStyle.Render(t.itemCount),
	)
}
