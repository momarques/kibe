package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/kube"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
	windowutil "github.com/momarques/kibe/internal/ui/window_util"
)

const tableViewHeightPercentage int = 32

type tableContent struct {
	syncState

	columns   []table.Column
	rows      []table.Row
	paginator paginatorModel
}

func newTableContent() tableContent {
	return tableContent{
		syncState: unsynced,
		paginator: newPaginatorModel(15),
	}
}

type tableModel struct {
	tableContent
	tableKeyMap
	table.Model

	response chan kube.TableResponse
}

func newTableModel() tableModel {
	t := table.New(
		table.WithFocused(true),
	)

	t.SetStyles(uistyles.NewTableStyle(false))
	t.SetHeight(
		windowutil.ComputeHeightPercentage(tableViewHeightPercentage))

	return tableModel{
		Model: t,

		response:     make(chan kube.TableResponse),
		tableContent: newTableContent(),
		tableKeyMap:  newTableKeyMap(),
	}
}

func (m tableModel) applyTableItems() (tableModel, tea.Cmd) {
	m.SetColumns(m.columns)

	start, end := m.paginator.GetSliceBounds(len(m.rows))
	m.SetRows(m.rows[start:end])
	return m, m.updateHeader(len(m.rows))
}

func (m CoreUI) updateOnTableResponse() (CoreUI, tea.Cmd) {
	var cmd tea.Cmd

	if response, ok := <-m.table.response; ok {
		m.table.rows = response.Rows
		m.table.columns = response.Columns

		m.table.paginator.SetTotalPages(len(m.table.rows))

		m.table, cmd = m.table.applyTableItems()

		m = m.changeSyncState(inSync)
		return m.updateStatusLog(m.logProcessDuration("OK", response.FetchDuration)),
			cmd
	}

	return m, nil
}

func (m CoreUI) updateTable(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch m.table.syncState {
	case inSync, syncing:
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:

			m.table.SetHeight(msg.Height - m.table.Height())
			// m.table.SetWidth(msg.Width - m.table.Width())
			m.table.SetColumns(m.client.ResourceSelected.Columns())
			m.help.Width = 20
			m.table.Model, cmd = m.table.Update(msg)
			return m, cmd

		case tea.KeyMsg:
			switch {
			case key.Matches(msg, m.table.Quit):
				return m, tea.Quit

			case key.Matches(msg, m.table.Describe):
				selectedResource := m.table.SelectedRow()

				m.tab, cmd = m.tab.describeResource(m.client, selectedResource[0])
				return m, cmd

			case key.Matches(msg, m.table.PreviousPage, m.table.NextPage):
				m.table.paginator.Model, _ = m.table.paginator.Update(msg)
				m.table, cmd = m.table.applyTableItems()

				return m, cmd
			}

		case descriptionReady:
			m.viewState = showTab
			m.tab.Tabs, m.tab.TabContent = msg.tabNames, msg.tabContent
			return m, nil

		case headerItemCountUpdated:
			m.header.itemCount = msg
			return m, nil

		case syncStarted:
			m, cmd = m.updateOnTableResponse()
			cmds = append(cmds, cmd)
			m, cmd = m.syncTable()
			cmds = append(cmds, cmd)

			return m, tea.Batch(cmds...)

		case spinner.TickMsg:
			m.syncBar.spinner, cmd = m.syncBar.spinner.Update(msg)
			return m, cmd
		}

	case unsynced:
		return m.syncTable()
	}

	m.table.Model, cmd = m.table.Update(msg)
	return m, cmd
}

func (m CoreUI) tableView() string {
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
