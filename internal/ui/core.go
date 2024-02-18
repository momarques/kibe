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
	showTab
)

type CoreUI struct {
	state  coreState
	height int

	client *kube.ClientReady

	listSelector *selector
	listUI       list.Model

	tableContent *content
	tableUI      table.Model

	tabUI tabModel

	headerUI    headerModel
	spinner     spinner.Model
	statusbarUI statusbar.Model
}

func NewUI() CoreUI {
	sp := spinner.New(
		spinner.WithStyle(uistyles.OKStatusMessage),
	)
	sp.Spinner = spinner.Dot

	status := newStatusBarUI()
	status.SetContent("Resource", "", "", "")

	selector := newListSelector(sp, status)
	content := newTableContent(nil)

	list := newListUI(selector)

	return CoreUI{
		state: showList,

		listSelector: selector,
		listUI:       list,

		tableContent: content,
		tableUI:      newTableUI(),

		tabUI: newTabUI(),

		spinner: sp,

		statusbarUI: status,
	}
}

func (m CoreUI) Init() tea.Cmd {
	return tea.SetWindowTitle("Kibe UI")
}

func (m CoreUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.QuitMsg:
		return m, tea.Quit
	}

	switch m.state {
	case showList:
		return m.updateListUI(msg)
	case showTable:
		return m.updateTableUI(msg)
	case showTab:
		return m.updateTabUI(msg)
	}
	return nil, nil
}

func (m CoreUI) View() string {
	switch m.state {

	case showList:
		return m.viewListUI()

	case showTable:
		return lipgloss.JoinVertical(
			lipgloss.Left,
			m.headerUI.viewHeaderUI(0),
			m.viewTableUI(),
			m.statusbarUI.View())
	case showTab:
		return lipgloss.JoinVertical(
			lipgloss.Left,
			m.headerUI.viewHeaderUI(0),
			m.viewTableUI(),
			m.viewTabUI(),
			m.statusbarUI.View())
	}
	return m.View()
}
