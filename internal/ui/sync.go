package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

	m.syncState = syncing
	m.syncBarModel = m.changeSyncState()

	m.tableContent.columns, m.tableContent.rows = m.client.FetchTableView()
	m.tableContent.paginatorModel.SetTotalPages(len(m.tableContent.rows))

	m.paginatorModel, _ = m.paginatorModel.Update(msg)
	m.tableModel, cmd = m.applyTableItems(m.tableModel)

	return m, tea.Batch(cmd, m.syncBarModel.spinner.Tick, func() tea.Msg {
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

func newSyncBarModel() syncBarModel {
	sp := spinner.New(
		spinner.WithStyle(uistyles.OKStatusMessage),
	)
	sp.Spinner = spinner.Dot
	return syncBarModel{
		spinner: sp,
	}
}

func (m CoreUI) changeSyncState() syncBarModel {
	switch m.syncState {
	case synced:
		m.syncBarModel.text = syncedText
		m.syncBarModel.color = lipgloss.Color(syncedColor)
		m.spinnerState = hideSpinner
	case syncing:
		m.syncBarModel.text = syncingText
		m.syncBarModel.color = lipgloss.Color(syncingColor)
		m.spinnerState = showSpinner
	case unsynced:
		m.syncBarModel.text = unsyncedText
		m.syncBarModel.color = lipgloss.Color(unsyncedColor)
		m.spinnerState = hideSpinner
	}
	return m.syncBarModel
}

func (m CoreUI) syncBarModelView() string {
	syncStyle := uistyles.ViewTitleStyle.
		Copy().
		MarginTop(1).
		MarginBottom(1)

	if m.spinnerState == showSpinner {
		return syncStyle.
			Background(m.syncBarModel.color).
			Render(m.spinner.View(), m.syncBarModel.text)
	}
	return syncStyle.
		Background(m.syncBarModel.color).
		Render(m.syncBarModel.text)
}
