package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/kube"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

type syncState int

const (
	inSync syncState = iota
	unsynced
	syncing

	inSyncColor   string = "#a4c847"
	unsyncedColor string = "#d83f24"
)

type syncStarted time.Time

func (s syncStarted) Cmd() func() tea.Msg {
	return func() tea.Msg {
		return s
	}
}

type syncBarModel struct {
	spinnerState

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
		spinner:      sp,
		spinnerState: hideSpinner,
	}
}

func (m CoreUI) changeSyncState(state syncState) CoreUI {
	m.table.syncState = state

	switch m.table.syncState {
	case inSync:
		m.syncBar.text = "in sync"
		m.syncBar.color = lipgloss.Color(inSyncColor)
		m.syncBar.spinnerState = showSpinner
	case unsynced:
		m.syncBar.text = "unsynced"
		m.syncBar.color = lipgloss.Color(unsyncedColor)
		m.syncBar.spinnerState = hideSpinner
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

func (m CoreUI) syncBarView() string {
	syncStyle := uistyles.ViewTitleStyle.
		Copy().
		Background(m.syncBar.color)

	if m.syncBar.spinnerState == showSpinner {
		return lipgloss.JoinHorizontal(lipgloss.Left,
			m.syncBar.spinner.View(),
			syncStyle.
				Render(m.syncBar.text))
	}
	return syncStyle.
		Render(m.syncBar.text)
}
