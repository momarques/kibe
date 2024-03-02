package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	m.syncState = syncing

	m.tableModel, cmd = m.fetchTableItems(m.tableModel)
	m.paginatorModel, _ = m.paginatorModel.Update(msg)
	m.tableModel = m.fetchPageItems(m.tableModel)

	return m, tea.Batch(cmd, func() tea.Msg {
		return lastSync(time.Now())
	})
}

func startSyncing(t time.Time) tea.Msg {
	return unsynced
}

type syncIndicatorModel struct {
	spinner spinner.Model
	text    string
}

func (m CoreUI) viewsyncIndicatorModel() string {
	switch m.tableContent.syncState {
	case synced:
		return lipgloss.NewStyle().
			Border(lipgloss.BlockBorder(), true).
			Background(lipgloss.Color("#ffffff")).Render("synced")

	}
	return ""
}
