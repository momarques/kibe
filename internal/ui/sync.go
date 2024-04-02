package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/logging"
	"github.com/momarques/kibe/internal/ui/style"
)

type syncState int

const (
	inSync syncState = iota
	notSynced
	syncing
	starting
	paused
)

type syncStarted time.Time

func (s syncStarted) Cmd() func() tea.Msg {
	return func() tea.Msg {
		return s
	}
}

type syncFinished time.Time

type syncBarModel struct {
	spinnerState

	bgColor lipgloss.Color
	fgColor lipgloss.Color
	spinner spinner.Model
	text    string
}

func newSyncBarModel() syncBarModel {
	return syncBarModel{
		fgColor: lipgloss.Color("#ffffff"),
		spinner: spinner.New(
			spinner.WithSpinner(spinner.Dot),
			spinner.WithStyle(style.OKStatusMessage()),
		),
		spinnerState: hideSpinner,
	}
}

func (m CoreUI) changeSyncState(state syncState) CoreUI {
	logging.Log.
		WithField("state", state).
		Debug("changing state")
	m.table.syncState = state

	switch m.table.syncState {
	case inSync:
		m.syncBar.text = "in sync"
		m.syncBar.bgColor = style.InSyncColor()
		m.syncBar.spinnerState = showSpinner
	case notSynced:
		m.syncBar.text = "not synced"
		m.syncBar.bgColor = style.NotSyncedColor()
		m.syncBar.spinnerState = hideSpinner
	case starting:
		m.syncBar.text = "starting"
		m.syncBar.bgColor = style.StartingColor()
		m.syncBar.spinnerState = showSpinner
	case paused:
		m.syncBar.text = "paused"
		m.syncBar.bgColor = style.PausedColor()
		m.syncBar.spinnerState = hideSpinner
	}
	return m
}

func (m CoreUI) syncTable() (CoreUI, tea.Cmd) {
	go func() {
		m.client.FetchTableViewAsync(m.table.response)
	}()

	m = m.changeSyncState(syncing)

	return m, tea.Batch(
		m.logProcess(m.client.LogOperation()),
		m.syncBar.spinner.Tick,
		syncStarted(time.Now()).Cmd(),
	)
}

func (s syncBarModel) spinnerView() string {
	if s.spinnerState == showSpinner {
		return s.spinner.View()
	}
	return ""
}

func (m CoreUI) syncBarView() string {
	syncStyle := style.SyncBarStatusStyle().
		Foreground(m.syncBar.fgColor).
		Background(m.syncBar.bgColor)

	return lipgloss.JoinHorizontal(lipgloss.Left,
		m.syncBar.spinnerView(),
		syncStyle.
			Render(m.syncBar.text))
}
