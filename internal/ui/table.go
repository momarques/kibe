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

func newTableModel() table.Model {
	t := table.New(
		table.WithFocused(true),
	)

	t.SetStyles(uistyles.NewTableStyle(false))
	t.SetHeight(
		windowutil.ComputePercentage(
			windowHeight, tableViewHeightPercentage))
	return t
}

func (m CoreUI) updatetableModel(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.tableContent.syncState {
	case synced, syncing:
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:

			// m.tableModel.SetHeight(msg.Height - m.tableModel.Height())
			// m.tableModel.SetWidth(msg.Width - m.tableModel.Width())
			logging.Log.Infof("window size -> %d x %d", msg.Width, msg.Height)
			m.helpModel.Width = 20
			m.tableModel, cmd = m.tableModel.Update(msg)
			return m, cmd

		case tea.KeyMsg:
			switch {
			case key.Matches(msg, m.tableKeyMap.Quit):
				return m, tea.Quit
			case key.Matches(msg, m.tableKeyMap.Describe):
				selectedResource := m.tableModel.SelectedRow()
				m.tabModel.Tabs, m.tabModel.TabContent = m.tabModel.describeResource(m.client, selectedResource[0])

				m.viewState = showTab
				return m, nil
			case key.Matches(msg, m.tableKeyMap.PreviousPage, m.tableKeyMap.NextPage):
				m.tableContent.paginator, _ = m.tableContent.paginator.Update(msg)
				m.tableModel = m.tableContent.fetchPageItems(m.tableModel)

				return m, cmd
			}
		case headerUpdated:
			m.headerModel.text = msg.text
			m.headerModel.itemCount = msg.itemCount
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

	m.tableModel, cmd = m.tableModel.Update(msg)
	return m, cmd
}

func (m CoreUI) viewtableModel() string {
	tableStyle := uistyles.TableStyle

	if m.viewState == showTab {
		tableStyle = uistyles.DimmedTableStyle
		m.tableModel.SetStyles(uistyles.NewTableStyle(true))
		return tableStyle.
			Render(m.tableModel.View())
	}
	return tableStyle.Render(m.tableModel.View())
}
