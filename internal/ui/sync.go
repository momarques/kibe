package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/logging"
)

type syncState int

const (
	synced syncState = iota
	unsynced
	syncing
)

type lastSync time.Time

func (m CoreUI) sync(msg tea.Msg) (CoreUI, tea.Cmd) {
	var cmd tea.Cmd

	m.tableContent.client = m.client
	m.tableContent.syncState = syncing

	logging.Log.Info("syncing...")

	m.tableUI, cmd = m.tableContent.fetchTableItems(m.tableUI)
	m.tableContent.paginator, _ = m.tableContent.paginator.Update(msg)
	m.tableUI = m.tableContent.fetchPageItems(m.tableUI)

	return m, tea.Batch(cmd, func() tea.Msg {
		return lastSync(time.Now())
	})
}

func startSyncing(t time.Time) tea.Msg {
	return unsynced
}

type syncIndicatorUI struct {
	state   syncState
	spinner spinner.Model
	text    string
}

func (m CoreUI) updateSyncIndicatorUI(msg tea.Msg) (tea.Model, tea.Cmd) {

	return m, nil
}

func (m CoreUI) viewSyncIndicatorUI() string {
	return ""
}
