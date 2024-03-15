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
	var logMsg string
	var fetchDuration time.Duration

	m.table.syncState = syncing
	m.syncBarModel = m.changeSyncState()

	now := time.Now()
	fetchDuration = func() time.Duration {

		m.table.columns, m.table.rows, logMsg = m.client.FetchTableView()
		m.table.paginator.SetTotalPages(len(m.table.rows))

		m.table.paginator, _ = m.table.paginator.Update(msg)
		m.table, cmd = m.table.applyTableItems()

		return time.Since(now)
	}()

	return m, tea.Batch(
		cmd,
		m.syncBarModel.spinner.Tick,
		m.logProcess(logMsg, fetchDuration),
		func() tea.Msg {
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
	switch m.table.syncState {
	case synced:
		m.syncBarModel.text = syncedText
		m.syncBarModel.color = lipgloss.Color(syncedColor)
		m.list.spinnerState = hideSpinner
	case syncing:
		m.syncBarModel.text = syncingText
		m.syncBarModel.color = lipgloss.Color(syncingColor)
		m.list.spinnerState = showSpinner
	case unsynced:
		m.syncBarModel.text = unsyncedText
		m.syncBarModel.color = lipgloss.Color(unsyncedColor)
		m.list.spinnerState = hideSpinner
	}
	return m.syncBarModel
}

func (m CoreUI) syncBarModelView() string {
	syncStyle := uistyles.ViewTitleStyle.
		Copy()

	if m.list.spinnerState == showSpinner {
		return syncStyle.
			Background(m.syncBarModel.color).
			Render(m.list.spinner.View(), m.syncBarModel.text)
	}
	return syncStyle.
		Background(m.syncBarModel.color).
		Render(m.syncBarModel.text)
}
