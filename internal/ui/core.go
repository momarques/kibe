package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mistakenelf/teacup/statusbar"
	"github.com/momarques/kibe/internal/kube"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
	windowutil "github.com/momarques/kibe/internal/ui/window_util"
)

var (
	windowWidth, windowHeight = windowutil.GetWindowSize()
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

	// main UIs
	listSelector *selector
	listUI       list.Model

	tableContent *content
	tableKeys    tableKeyMap
	tableUI      table.Model

	tabKeys tabKeyMap
	tabUI   tabModel

	// utility UIs
	helpUI    help.Model
	headerUI  headerModel
	spinner   spinner.Model
	statusbar statusbar.Model
}

func NewUI() CoreUI {
	sp := spinner.New(
		spinner.WithStyle(uistyles.OKStatusMessage),
	)
	sp.Spinner = spinner.Dot

	status := newStatusBar()
	status.SetContent("Resource", "", "", "")

	selector := newListSelector(sp, status)
	paginator := newPaginatorUI()

	content := newTableContent(nil, paginator)

	list := newListUI(selector)

	return CoreUI{
		state: showList,

		listSelector: selector,
		listUI:       list,

		tableContent: content,
		tableKeys:    tableShortcuts,
		tableUI:      newTableUI(),

		tabKeys: tabShortcuts,
		tabUI:   newTabUI(),

		helpUI:    help.New(),
		spinner:   sp,
		statusbar: status,
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

	case showTable, showTab:
		return m.viewMainUI()
	}
	return m.View()
}

func (m CoreUI) viewMainUI() string {
	var helpBindingLines [][]key.Binding

	switch m.state {
	case showTable:
		helpBindingLines = append(helpBindingLines,
			m.tableKeys.viewFirstLine(),
			m.tableKeys.viewSecondLine())

	case showTab:
		helpBindingLines = append(helpBindingLines,
			m.tabKeys.viewFirstLine())
	}

	helpView := lipgloss.JoinVertical(
		lipgloss.Center,
		m.showHelp(helpBindingLines...)...)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.headerUI.viewHeaderUI(0),
		m.viewTableUI(),
		m.viewTabUI(),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.viewPaginatorUI(),
			helpView,
		),
		m.statusbar.View())
}

func (m CoreUI) showHelp(helpBindingLines ...[]key.Binding) []string {
	var helpLines []string

	helpStyle := lipgloss.NewStyle().MarginBottom(1)

	for _, line := range helpBindingLines {
		helpLines = append(helpLines, helpStyle.Render(
			m.helpUI.ShortHelpView(line)))
	}
	return helpLines
}
