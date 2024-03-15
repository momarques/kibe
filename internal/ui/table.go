package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/kube"
	"github.com/momarques/kibe/internal/logging"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
	windowutil "github.com/momarques/kibe/internal/ui/window_util"
)

const tableViewHeightPercentage int = 32

type tableModel struct {
	tableContent
	table.Model
}

func newTableModel() tableModel {
	t := table.New(
		table.WithFocused(true),
	)

	t.SetStyles(uistyles.NewTableStyle(false))
	t.SetHeight(
		windowutil.ComputeHeightPercentage(tableViewHeightPercentage))

	return tableModel{
		tableContent: newTableContent(),
		Model:        t,
	}
}

func (m CoreUI) updateTableModel(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.table.syncState {
	case synced, syncing:
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:

			m.table.SetHeight(msg.Height - m.table.Height())
			// m.table.SetWidth(msg.Width - m.table.Width())
			m.table.SetColumns(m.client.ResourceSelected.Columns())
			logging.Log.Infof("window size -> %d x %d", msg.Width, msg.Height)
			logging.Log.Infof("table size -> %d x %d", m.table.Width(), m.table.Height())
			m.helpModel.Width = 20
			m.table.Model, cmd = m.table.Update(msg)
			return m, cmd

		case tea.KeyMsg:
			switch {
			case key.Matches(msg, m.tableKeys.Quit):
				return m, tea.Quit

			case key.Matches(msg, m.tableKeys.Describe):
				selectedResource := m.table.SelectedRow()

				m.tab, cmd = m.tab.describeResource(m.client, selectedResource[0])
				return m, cmd

			case key.Matches(msg, m.tableKeys.PreviousPage, m.tableKeys.NextPage):
				m.table.paginator, _ = m.table.paginator.Update(msg)
				m.table, cmd = m.table.applyTableItems()

				return m, cmd
			}

		case descriptionReady:
			m.viewState = showTab
			m.tab.Tabs, m.tab.TabContent = msg.tabNames, msg.tabContent
			return m, nil

		case headerItemCountUpdated:
			m.headerModel.itemCount = msg
			return m, nil

		case lastSync:
			m.table.syncState = synced
			m.syncBarModel = m.changeSyncState()

			return m, tea.Batch(tea.Tick(kube.ResquestTimeout, startSyncing))

		case syncState:
			if msg == unsynced {
				m.table.syncState = msg
				m.syncBarModel = m.changeSyncState()
				return m.sync(nil)
			}
		}

	case unsynced:
		return m.sync(msg)
	}

	m.table.Model, cmd = m.table.Update(msg)
	return m, cmd
}

func (m CoreUI) tableModelView() string {
	tableStyle := uistyles.TableStyle

	if m.viewState == showTab {
		tableStyle = uistyles.DimmedTableStyle
		m.table.SetStyles(uistyles.NewTableStyle(true))
		return tableStyle.Render(m.table.View())
	}
	if m.table.columns == nil {
		return lipgloss.NewStyle().
			Height((windowutil.ComputeHeightPercentage(tableViewHeightPercentage) + 3)).
			Render("")
	}
	return tableStyle.Render(m.table.View())
}
