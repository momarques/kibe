package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
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

	keys         enabledKeys
	list         list.Model
	listSelector *listSelector
	table        tableModel
	tab          tabModel
	tabKeys      tabKeyMap

	headerModel    headerModel
	helpModel      help.Model
	statusbarModel statusbar.Model
	syncBarModel   syncBarModel
	statusLogModel
}

func NewUI() CoreUI {
	selector := newListSelector()

	tabKeyMap := newTabKeyMap()
	table := newTableModel()

	return CoreUI{
		viewState: showList,

		keys:         setKeys(table.tableKeyMap, tabKeyMap),
		list:         newlistModel(selector),
		listSelector: selector,
		table:        table,
		tab:          newTabModel(),
		tabKeys:      tabKeyMap,

		headerModel:    headerModel{},
		helpModel:      help.New(),
		statusbarModel: newStatusBarModel(),
		statusLogModel: newStatusLogModel(),
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
	case statusLogMessage:
		return m.updateStatusLog(msg), nil
	}

	switch m.viewState {
	case showList:
		return m.updateListModel(msg)
	case showTable:
		m.keys = m.keys.setEnabled(m.table.fullHelp()...)
		return m.updateTableModel(msg)
	case showTab:
		switch m.tab.tabViewState {
		case contentSelected:
			m.keys = m.keys.setEnabled(m.tabKeys.fullHelp()...)
		case noContentSelected:
			m.keys = m.keys.setEnabled(m.tabKeys.fullHelpWithContentSelected()...)
		}
		return m.updateTabModel(msg)
	}
	return m, nil
}

func (m CoreUI) View() string {
	switch m.viewState {

	case showList:
		return m.listView()

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
		helpBindingLines = [][]key.Binding{
			m.table.firstHelpLineView(),
			m.table.secondHelpLineView(),
		}

	case showTab:
		switch m.tab.tabViewState {
		case noContentSelected:
			helpBindingLines = [][]key.Binding{
				m.tabKeys.firstHelpLineView(),
				m.tabKeys.secondHelpLineView(),
			}
		case contentSelected:
			helpBindingLines = [][]key.Binding{
				m.tabKeys.firstHelpLineViewWithContentSelected(),
				m.tabKeys.secondHelpLineView(),
			}
		}
	}

	helpView := lipgloss.JoinVertical(
		lipgloss.Center,
		m.showHelpLines(helpBindingLines...)...)

	leftUtilityPanel := lipgloss.JoinVertical(
		lipgloss.Left,
		m.paginatorModelView(),
		m.syncBarModelView(),
	)

	bottomPanel := lipgloss.JoinVertical(lipgloss.Left,
		m.tabView(),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			leftUtilityPanel,
			helpView,
		))

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.headerModelView(),
		m.tableModelView(),
		lipgloss.JoinHorizontal(lipgloss.Center,
			bottomPanel,
			m.statusLogModelView()),
		m.statusbarModel.View())
}
