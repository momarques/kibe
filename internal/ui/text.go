package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

type headerModel struct {
	text, line string
}

type headerUpdated struct {
	text, line string
}

func (c *content) updateHeader(title string) tea.Cmd {
	return func() tea.Msg {
		return headerUpdated{text: title, line: ""}
	}
}

func (t headerModel) viewHeaderUI(size int) string {
	textStyle := uistyles.
		ViewTitleStyle.
		Copy().
		PaddingLeft(1).
		MarginLeft(2).
		MarginBottom(1)

	lineStyle := textStyle.
		Copy().
		MarginBottom(4).
		Width(size)
		// .
		// Border(lipgloss.ThickBorder(), false, false, true, false).
		// BorderBottomForeground(lipgloss.Color("#d65f50"))

	return lipgloss.JoinVertical(lipgloss.Top,
		textStyle.Render(t.text),
		lineStyle.Render(t.line),
	)
}
