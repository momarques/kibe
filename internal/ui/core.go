package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mistakenelf/teacup/statusbar"
	"github.com/momarques/kibe/internal/kube"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

type coreState int

const (
	showList coreState = iota
	showTable
)

type CoreUI struct {
	state  coreState
	height int

	client *kube.ClientReady

	listSelector *selector
	listUI       list.Model

	tableContent *content
	tableUI      table.Model

	spinner spinner.Model

	statusbarUI *statusbar.Model
}

func NewUI() CoreUI {
	sp := spinner.New(
		spinner.WithStyle(uistyles.OKStatusMessage),
	)
	sp.Spinner = spinner.Dot

	status := newStatusBarUI()
	status.SetContent("Resource", "", "", "")

	selector := newListSelector(sp, &status)
	content := newTableContent(nil)

	list := newListUI(selector)
	return CoreUI{
		state: showList,

		listSelector: selector,
		listUI:       list,

		tableContent: content,
		tableUI:      newTableUI(),

		spinner: sp,

		statusbarUI: &status,
	}
}

func (m CoreUI) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m CoreUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case showList:
		return m.updateListUI(msg)
	case showTable:
		return m.updateTableUI(msg)
	}
	return nil, nil
}

func (m CoreUI) View() string {
	switch m.state {

	case showList:
		return m.viewListUI()

	case showTable:
		return m.viewTableUI()
	}
	return lipgloss.JoinVertical(
		lipgloss.Top, m.View(),
		m.statusbarUI.View())
}
