package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/kube"
	"github.com/momarques/kibe/internal/ui/style"
)

type syncState int

const (
	inSync syncState = iota
	unsynced
	syncing
	starting
)

type syncStarted time.Time

func (s syncStarted) Cmd() func() tea.Msg {
	return func() tea.Msg {
		return s
	}
}

type syncBarModel struct {
	spinnerState

	bgColor lipgloss.Color
	fgColor lipgloss.Color
	spinner spinner.Model
	text    string
}

func newSyncBarModel() syncBarModel {
	sp := spinner.New(
		spinner.WithStyle(style.OKStatusMessage()),
	)
	sp.Spinner = spinner.Dot
	return syncBarModel{
		fgColor:      lipgloss.Color("#ffffff"),
		spinner:      sp,
		spinnerState: hideSpinner,
	}
}

func (m CoreUI) changeSyncState(state syncState) CoreUI {
	m.table.syncState = state

	switch m.table.syncState {
	case inSync:
		m.syncBar.text = "in sync"
		m.syncBar.bgColor = style.InSyncColor()
		m.syncBar.spinnerState = showSpinner
	case unsynced:
		m.syncBar.text = "unsynced"
		m.syncBar.bgColor = style.UnsyncedColor()
		m.syncBar.spinnerState = hideSpinner
	case starting:
		m.syncBar.text = "starting"
		m.syncBar.bgColor = style.StartingColor()
		m.syncBar.spinnerState = showSpinner
	}
	return m
}

func (m CoreUI) syncTable() (CoreUI, tea.Cmd) {
	m = m.changeSyncState(syncing)

	go func() {
		m.client.FetchTableView(m.table.response)
	}()

	return m, tea.Batch(
		m.logProcess(m.client.LogOperation()),
		m.syncBar.spinner.Tick,
		tea.Tick(kube.ResquestTimeout, func(t time.Time) tea.Msg {
			return syncStarted(time.Now())
		}),
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
