package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

func (s syncState) String() string {
	switch s {
	case inSync:
		return "in sync"
	case notSynced:
		return "not synced"
	case starting:
		return "starting"
	case paused:
		return "paused"
	case syncing:
		return "syncing"
	default:
		return ""
	}
}

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
			spinner.WithStyle(style.ClientConfigSpinnerStyle()),
		),
		spinnerState: hideSpinner,
	}
}

func (m CoreUI) changeSyncState(state syncState) CoreUI {
	m.log.WithDebugContext(m.client).
		Str("state", m.table.syncState.String()).
		Msg("changing table sync state")
	m.table.syncState = state

	switch m.table.syncState {
	case inSync:
		m.syncBar.text = m.table.syncState.String()
		m.syncBar.bgColor = style.InSyncColor()
		m.syncBar.spinnerState = showSpinner
	case notSynced:
		m.syncBar.text = m.table.syncState.String()
		m.syncBar.bgColor = style.NotSyncedColor()
		m.syncBar.spinnerState = hideSpinner
	case starting:
		m.syncBar.text = m.table.syncState.String()
		m.syncBar.bgColor = style.StartingColor()
		m.syncBar.spinnerState = showSpinner
	case paused:
		m.syncBar.text = m.table.syncState.String()
		m.syncBar.bgColor = style.PausedColor()
		m.syncBar.spinnerState = hideSpinner
	}
	return m
}

func (m CoreUI) syncTable() (CoreUI, tea.Cmd) {
	m.log.WithDebugContext(m.client).Msg("will fetch table view asynchronously")
	go func() {
		m.client.FetchTableViewAsync(m.table.response)
	}()

	m = m.changeSyncState(syncing)
	m.log.WithDebugContext(m.client).Msg(m.client.LogOperation())

	return m, tea.Batch(
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
