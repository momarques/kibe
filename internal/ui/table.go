package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/kube"
	"github.com/momarques/kibe/internal/logging"
	"github.com/momarques/kibe/internal/ui/style"
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
		syncState: starting,
		paginator: newPaginatorModel((windowutil.ComputeHeightPercentage(tableViewHeightPercentage))),
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

	t.SetStyles(style.NewTableStyle(false))
	t.SetHeight(
		windowutil.ComputeHeightPercentage(tableViewHeightPercentage))

	return tableModel{
		Model: t,

		response:     make(chan kube.TableResponse),
		tableContent: newTableContent(),
		tableKeyMap:  newTableKeyMap(),
	}
}

func (m tableModel) applyTableItems(r kube.TableResponse) (tableModel, tea.Cmd) {
	m.rows = r.Rows
	m.columns = r.Columns
	m.SetColumns(m.columns)

	m.paginator.SetTotalPages(len(m.rows))
	return m, m.updateHeader(len(m.rows))
}

func (m tableModel) applyPageChanges() tableModel {
	start, end := m.paginator.GetSliceBounds(len(m.rows))
	m.SetRows(m.rows[start:end])
	return m
}

func (m CoreUI) updateOnTableResponse() (CoreUI, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	response := m.client.FetchTableView()
	m.table, cmd = m.table.applyTableItems(response)
	cmds = append(cmds, cmd)

	m.table = m.table.applyPageChanges()

	m = m.changeSyncState(inSync)
	m, cmd = m.syncTable()
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m CoreUI) checkTableResponseAsync() (CoreUI, tea.Cmd) {
	if response, ok := <-m.table.response; ok {
		return m.updateTableWithAsyncResponse(response)
	}
	return m, nil
}

func (m CoreUI) updateTableWithAsyncResponse(response kube.TableResponse) (CoreUI, tea.Cmd) {
	var cmd tea.Cmd

	if response.Err != nil {
		m = m.changeSyncState(notSynced)
		return m.updateStatusLog(m.logProcessDuration(NOK, response.FetchDuration)), nil
	}

	m.table, cmd = m.table.applyTableItems(response)
	m.table = m.table.applyPageChanges()

	m = m.changeSyncState(inSync)
	return m.updateStatusLog(m.logProcessDuration(OK, response.FetchDuration)),
		cmd
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
				m.client.ResourceSelected = m.client.SetID(m.table.SelectedRow()[0])

				m.tab, cmd = m.tab.describeResource(m.client)

				return m, tea.Batch(cmd,
					updateStatusBar(fmt.Sprintf("%s: %s", m.client.Kind(), m.client.ID()),
						m.client.ContextSelected.String(),
						m.client.NamespaceSelected.String()))

			case key.Matches(msg, m.table.PreviousPage, m.table.NextPage):
				m.table.paginator.Model, _ = m.table.paginator.Update(msg)
				m.table = m.table.applyPageChanges()

				return m, nil
			}

		case descriptionReady:
			m.viewState = showTab
			m.tab.Tabs, m.tab.TabContent = msg.tabNames, msg.tabContent
			return m, nil

		case headerItemCountUpdated:
			m.header.itemCount = msg
			return m, nil

		case syncStarted:
			logging.Log.Info("syncStarted")
			return m, tea.Tick(kube.ResquestTimeout,
				func(t time.Time) tea.Msg {
					return syncFinished(time.Now())
				})

		case syncFinished:
			logging.Log.Info("syncFinished")

			m, cmd = m.checkTableResponseAsync()
			cmds = append(cmds, cmd)
			m, cmd = m.syncTable()
			cmds = append(cmds, cmd)

			return m, tea.Batch(cmds...)

		case spinner.TickMsg:
			m.syncBar.spinner, cmd = m.syncBar.spinner.Update(msg)
			return m, cmd
		}

	case notSynced:
		logging.Log.Info("not synced")
		return m.syncTable()

	case starting:
		logging.Log.Info("starting")
		return m.updateOnTableResponse()

	case paused:
		logging.Log.Info("paused")
		return m, nil
	}

	m.table.Model, cmd = m.table.Update(msg)
	return m, cmd
}

func (m CoreUI) tableView() string {
	tableStyle := style.TableStyle

	if m.viewState == showTab {
		tableStyle = style.DimmedTableStyle
		m.table.SetStyles(style.NewTableStyle(true))
		return tableStyle().Render(m.table.View())
	}
	if m.table.columns == nil {
		return lipgloss.NewStyle().
			Height((windowutil.ComputeHeightPercentage(tableViewHeightPercentage) + 3)).
			Render("")
	}
	return tableStyle().Render(m.table.View())
}
