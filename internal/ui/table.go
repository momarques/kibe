package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

func newTableUI() table.Model {
	t := table.New(
		table.WithFocused(true),
		table.WithHeight(50),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("24")).
		BorderBottom(true).
		BorderTop(true).
		Bold(true)

	// s.Cell = s.Cell.MarginRight(5)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("24")).
		Bold(false)

	t.SetStyles(s)

	return t
}

func (m CoreUI) updateTableUI(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.tableContent.contentState {
	case loaded:
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			h, v := uistyles.AppStyle.GetFrameSize()
			m.tableUI.SetHeight(h)
			m.tableUI.SetWidth(v)

		case tea.KeyMsg:

			switch msg.String() {
			case "esc":
				if m.tableUI.Focused() {
					m.tableUI.Blur()
				} else {
					m.tableUI.Focus()
				}
			case "q", "ctrl+c":
				return m, tea.Quit
			case "enter":
				return m, tea.Batch(
					tea.Printf("Let's go to %s!", m.tableUI.SelectedRow()[1]),
				)
			}
		default:
			return m, tea.Tick(loadInterval, func(t time.Time) tea.Msg {
				m.tableContent.contentState = notLoaded
				return nil
			})
		}

	case notLoaded:
		m.tableContent.client = m.client
		m.tableUI = *m.tableContent.fetch(&m.tableUI)
	}

	m.tableUI, cmd = m.tableUI.Update(msg)
	return m, cmd
}

func (m CoreUI) viewTableUI() string {
	return lipgloss.JoinVertical(lipgloss.Bottom,
		lipgloss.Place(1, 1, lipgloss.Center, lipgloss.Center, m.tableUI.View()),
		lipgloss.Place(
			1, 1,
			lipgloss.Center, lipgloss.Bottom,
			m.statusbarUI.View()),
	)
}
