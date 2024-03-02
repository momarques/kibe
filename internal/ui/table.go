package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/logging"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
	windowutil "github.com/momarques/kibe/internal/ui/window_util"
)

const tableViewHeightPercentage int = 30

func newTableUI() table.Model {
	t := table.New(
		table.WithFocused(true),
	)

	t.SetStyles(uistyles.NewTableStyle(false))
	t.SetHeight(
		windowutil.ComputePercentage(
			windowHeight, tableViewHeightPercentage))
	return t
}

func (m CoreUI) updateTableUI(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.tableContent.syncState {
	case synced, syncing:
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:

			// m.tableUI.SetHeight(msg.Height - m.tableUI.Height())
			// m.tableUI.SetWidth(msg.Width - m.tableUI.Width())
			logging.Log.Infof("window size -> %d x %d", msg.Width, msg.Height)
			m.helpUI.Width = 20
			m.tableUI, cmd = m.tableUI.Update(msg)
			return m, cmd

		case tea.KeyMsg:
			switch {
			case key.Matches(msg, m.tableKeys.Quit):
				return m, tea.Quit
			case key.Matches(msg, m.tableKeys.Describe):
				selectedResource := m.tableUI.SelectedRow()
				m.tabUI.Tabs, m.tabUI.TabContent = m.tabUI.describeResource(m.client, selectedResource[0])

				m.state = showTab
				return m, nil
			case key.Matches(msg, m.tableKeys.PreviousPage, m.tableKeys.NextPage):
				m.tableContent.paginator, _ = m.tableContent.paginator.Update(msg)
				m.tableUI = m.tableContent.fetchPageItems(m.tableUI)

				return m, cmd
			}
		case headerUpdated:
			m.headerUI.text = msg.text
			m.headerUI.itemCount = msg.itemCount
			return m, nil
		case lastSync:
			m.tableContent.syncState = synced

			return m, tea.Batch(tea.Tick(loadInterval, startSyncing))
		case syncState:
			if msg == unsynced {
				m.tableContent.syncState = unsynced
				return m.sync(nil)
			}
		}

	case unsynced:
		return m.sync(msg)
	}

	m.tableUI, cmd = m.tableUI.Update(msg)
	return m, cmd
}

func (m CoreUI) viewTableUI() string {
	tableStyle := uistyles.TableStyle

	if m.state == showTab {
		tableStyle = uistyles.DimmedTableStyle
		m.tableUI.SetStyles(uistyles.NewTableStyle(true))
		return tableStyle.
			Render(m.tableUI.View())
	}
	return tableStyle.Render(m.tableUI.View())
}
