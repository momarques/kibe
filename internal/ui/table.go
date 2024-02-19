package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/logging"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

const tableViewHeightPercentage int = 30

func newTableUI() table.Model {
	t := table.New(
		table.WithFocused(true),
	)

	s := table.DefaultStyles()

	s.Cell = uistyles.TableCellStyle.Copy()
	s.Header = uistyles.TableHeaderStyle.Copy()
	s.Selected = uistyles.TableSelectedStyle.Copy()

	// s.Header = s.Header.
	// 	Border(lipgloss.NormalBorder()).
	// 	BorderForeground(lipgloss.Color("99")).
	// 	Bold(true).
	// 	Foreground(lipgloss.Color("99"))

	// s.Cell = s.Cell.
	// 	Border(lipgloss.NormalBorder()).
	// 	BorderForeground(lipgloss.Color("99")).
	// 	Foreground(lipgloss.Color("229"))

	// s.Selected = s.Selected.
	// 	BorderForeground(lipgloss.Color("#ffffff")).
	// 	Foreground(lipgloss.Color("#ffffff"))

	t.SetStyles(s)
	return t
}

func (m CoreUI) updateTableUI(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.tableContent.contentState {
	case loaded:
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:

			// m.tableUI.SetHeight(msg.Height - m.tableUI.Height())
			// m.tableUI.SetWidth(msg.Width - m.tableUI.Width())
			logging.Log.Infof("window size -> %d x %d", msg.Width, msg.Height)

			m.tableUI, cmd = m.tableUI.Update(msg)
			return m, cmd

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
				selectedResource := m.tableUI.SelectedRow()
				m.tabUI.Tabs, m.tabUI.TabContent = m.tabUI.describeResource(m.client, selectedResource[0])

				m.state = showTab
				return m, nil
			}
		case headerUpdated:
			m.headerUI.text = msg.text
			m.headerUI.itemCount = msg.itemCount

		default:
			return m, tea.Tick(loadInterval, func(t time.Time) tea.Msg {
				m.tableContent.contentState = notLoaded
				return nil
			})
		}

	case notLoaded:
		m.tableContent.client = m.client

		m.tableUI, cmd = m.tableContent.fetch(m.tableUI)
		return m, cmd
	}

	m.tableUI, cmd = m.tableUI.Update(msg)
	return m, cmd
}

func (m CoreUI) viewTableUI() string {
	tableView := uistyles.TableStyle.
		Copy().
		MarginLeft(2).
		Border(lipgloss.DoubleBorder(), true, true, true, true).
		BorderForeground(lipgloss.Color("#ffb8bc")).
		// Height(1).
		Render(m.tableUI.View())

	return tableView
}
