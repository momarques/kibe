package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

type headerModel struct {
	text, line string
	itemCount  string
}

type headerUpdated struct {
	text, line string
	itemCount  string
}

func (c *content) updateHeader(title string, itemCount int) tea.Cmd {
	return func() tea.Msg {
		return headerUpdated{
			text:      title,
			line:      "",
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

	// lineStyle := titleStyle.
	// 	Copy().
	// 	MarginBottom(0).
	// 	MarginTop(0).
	// 	Width(windowWidth).
	// 	Border(lipgloss.DoubleBorder(), true, false, false, false).
	// 	BorderTopForeground(lipgloss.Color("#d65f50")).
	// 	Background(lipgloss.NoColor{})

	itemCountStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#685e59")).
		MarginTop(1).
		MarginLeft(2).
		MarginBottom(3)

	return lipgloss.JoinVertical(lipgloss.Left,
		titleStyle.Render(t.text),
		// lineStyle.Render(t.line),
		itemCountStyle.Render(t.itemCount),
	)
}
