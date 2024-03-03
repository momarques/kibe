package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/logging"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

type syncState int

const (
	synced syncState = iota
	unsynced
	syncing

	syncedText   string = "synced"
	syncingText  string = "syncing"
	unsyncedText string = "unsynced"

	syncedColor   string = "#a4c847"
	syncingColor  string = "#cea540"
	unsyncedColor string = "#d83f24"
)

type lastSync time.Time

func (m CoreUI) sync(msg tea.Msg) (CoreUI, tea.Cmd) {
	var cmd tea.Cmd

	m.tableContent.client = m.client
	m.syncState = syncing
	m.syncBarModel = m.changeSyncState()

	m.tableContent.columns, m.tableContent.rows = m.client.FetchTableView(windowWidth)
	m.tableContent.paginatorModel.SetTotalPages(len(m.tableContent.rows))

	m.paginatorModel, _ = m.paginatorModel.Update(msg)
	m.tableModel, cmd = m.applyTableItems(m.tableModel)

	return m, tea.Batch(cmd, func() tea.Msg {
		return lastSync(time.Now())
	})
}

func startSyncing(t time.Time) tea.Msg {
	return unsynced
}

type syncBarModel struct {
	spinner spinner.Model
	text    string
	color   lipgloss.Color
}

func (m CoreUI) changeSyncState() syncBarModel {
	logging.Log.Info("state ->> ", m.syncState)
	logging.Log.Info("text ->> ", m.syncBarModel.text)
	logging.Log.Info("color ->> ", m.syncBarModel.color)

	switch m.syncState {
	case synced:
		m.syncBarModel.text = syncedText
		m.syncBarModel.color = lipgloss.Color(syncedColor)
	case syncing:
		m.syncBarModel.text = syncingText
		m.syncBarModel.color = lipgloss.Color(syncingColor)
	case unsynced:
		m.syncBarModel.text = unsyncedText
		m.syncBarModel.color = lipgloss.Color(unsyncedColor)
	}
	return m.syncBarModel
}

func (m CoreUI) viewSyncBarModel() string {
	syncStyle := uistyles.ViewTitleStyle.
		Copy().
		MarginTop(1).
		MarginBottom(1)

	return syncStyle.
		Background(m.syncBarModel.color).
		Render(m.syncBarModel.text)
}
