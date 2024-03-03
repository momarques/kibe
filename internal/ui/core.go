package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mistakenelf/teacup/statusbar"
	"github.com/momarques/kibe/internal/kube"
)

type viewState int

const (
	showList viewState = iota
	showTable
	showTab
)

type CoreUI struct {
	viewState
	height int

	client *kube.ClientReady

	listModel list.Model
	*listSelector

	tableModel table.Model
	*tableContent
	tableKeyMap

	tabModel tabModel
	tabKeyMap

	headerModel    headerModel
	helpModel      help.Model
	statusbarModel statusbar.Model
	syncBarModel   syncBarModel
}

func NewUI() CoreUI {
	selector := newListSelector()

	return CoreUI{
		viewState: showList,

		listSelector: selector,
		listModel:    newlistModel(selector),

		tableContent: newTableContent(),
		tableKeyMap:  newTableKeyMap(),
		tableModel:   newTableModel(),

		tabKeyMap: newTabKeyMap(),
		tabModel:  newTabModel(),

		headerModel:    headerModel{},
		helpModel:      help.New(),
		statusbarModel: newStatusBarModel(),
		syncBarModel:   newSyncBarModel(),
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

	switch m.viewState {
	case showList:
		return m.updatelistModel(msg)
	case showTable:
		return m.updatetableModel(msg)
	case showTab:
		return m.updatetabModel(msg)
	}
	return m, nil
}

func (m CoreUI) View() string {
	switch m.viewState {

	case showList:
		return m.listModelView()

	case showTable, showTab:
		return m.composedView()
	}
	return m.View()
}

func (m CoreUI) showHelpLines(helpBindingLines ...[]key.Binding) []string {
	var helpLines []string

	helpStyle := lipgloss.NewStyle().MarginBottom(1)

	for _, line := range helpBindingLines {
		helpLines = append(helpLines, helpStyle.Render(
			m.helpModel.ShortHelpView(line)))
	}
	return helpLines
}

func (m CoreUI) composedView() string {
	var helpBindingLines [][]key.Binding

	switch m.viewState {
	case showTable:
		helpBindingLines = append(helpBindingLines,
			m.tableKeyMap.firstHelpLineView(),
			m.tableKeyMap.secondHelpLineView())

	case showTab:
		helpBindingLines = append(helpBindingLines,
			m.tabKeyMap.firstHelpLineView())
	}

	helpView := lipgloss.JoinVertical(
		lipgloss.Center,
		m.showHelpLines(helpBindingLines...)...)

	leftUtilityPanel := lipgloss.JoinVertical(
		lipgloss.Left,
		m.paginatorModelView(),
		m.syncBarModelView(),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.headerModelView(),
		m.tableModelView(),
		m.tabModelView(),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			leftUtilityPanel,
			helpView,
		),
		m.statusbarModel.View())
}
