package ui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/logging"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
	windowutil "github.com/momarques/kibe/internal/ui/window_util"
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

	t.SetStyles(s)
	t.SetHeight(
		windowutil.ComputePercentage(
			windowHeight, tableViewHeightPercentage))
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
			m.helpUI.Width = 20
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
			case "right", "left":
				m.tableContent.paginator, _ = m.tableContent.paginator.Update(msg)
				m.tableUI = m.tableContent.fetchPageItems(m.tableUI)

				return m, cmd
			}
		case headerUpdated:
			m.headerUI.text = msg.text
			m.headerUI.itemCount = msg.itemCount
			return m, nil
			/*tea.Tick(loadInterval, func(t time.Time) tea.Msg {
				m.tableContent.contentState = notLoaded
				return nil
			})*/
		}

	case notLoaded:
		m.tableContent.client = m.client

		// m.state = showTab
		m.tableUI, cmd = m.tableContent.fetchTableItems(m.tableUI)
		m.tableContent.paginator, _ = m.tableContent.paginator.Update(msg)
		m.tableUI = m.tableContent.fetchPageItems(m.tableUI)

		m.tableContent.contentState = loaded

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
