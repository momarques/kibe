package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/kube"
	"github.com/momarques/kibe/internal/logging"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

type syncState int

const (
	synced syncState = iota
	unsynced
	syncing

	syncedColor   string = "#a4c847"
	syncingColor  string = "#cea540"
	unsyncedColor string = "#d83f24"
)

type lastSync time.Time

func (l lastSync) Cmd() func() tea.Msg {
	return func() tea.Msg {
		return l
	}
}

func (m CoreUI) sync() (CoreUI, tea.Cmd) {
	var cmd tea.Cmd
	var logMsg string

	logging.Log.Info("sync -> ")
	m, cmd = m.changeSyncState(syncing)

	go func() {
		logMsg = m.client.FetchTableView(m.table.response)
	}()

	return m, tea.Batch(
		cmd,
		m.logProcess(logMsg),
		tea.Tick(kube.ResquestTimeout, func(t time.Time) tea.Msg {
			return syncStarted(time.Now())
		}),
	)
}

// func startSyncing(t time.Time) tea.Msg {
// 	return unsynced
// }

type syncStarted time.Time

func (s syncStarted) Cmd() func() tea.Msg {
	return func() tea.Msg {
		return s
	}
}

type syncBarModel struct {
	color   lipgloss.Color
	spinner spinner.Model
	text    string
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

func (m CoreUI) changeSyncState(state syncState) (CoreUI, tea.Cmd) {
	m.table.syncState = state

	logging.Log.Info("changing state to: ", state)
	switch state {
	case synced:
		m.syncBar.text = "synced"
		m.syncBar.color = lipgloss.Color(syncedColor)
		m.spinnerState = hideSpinner
	case syncing:
		m.syncBar.text = "syncing"
		m.syncBar.color = lipgloss.Color(syncingColor)
		m.spinnerState = showSpinner
	case unsynced:
		m.syncBar.text = "unsynced"
		m.syncBar.color = lipgloss.Color(unsyncedColor)
		m.spinnerState = hideSpinner
	}
	return m, m.syncBar.spinner.Tick
}

func (m CoreUI) syncBarView() string {
	syncStyle := uistyles.ViewTitleStyle.
		Copy()

	if m.list.spinnerState == showSpinner {
		return syncStyle.
			Background(m.syncBar.color).
			Render(m.list.spinner.View(), m.syncBar.text)
	}
	return syncStyle.
		Background(m.syncBar.color).
		Render(m.syncBar.text)
}
