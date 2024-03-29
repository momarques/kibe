package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mistakenelf/teacup/statusbar"
	"github.com/momarques/kibe/internal/kube"
	windowutil "github.com/momarques/kibe/internal/ui/window_util"
)

const blankSpaceHeightPercentage int = 3

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

	keys  enabledKeys
	list  listModel
	tab   tabModel
	table tableModel

	header    headerModel
	help      help.Model
	statusBar statusbar.Model
	statusLog statusLogModel
	syncBar   syncBarModel
}

func NewUI() CoreUI {
	tab := newTabModel()
	table := newTableModel()

	return CoreUI{
		viewState: showList,

		keys:  setKeys(table.tableKeyMap, tab.tabKeyMap),
		list:  newListModel(),
		tab:   tab,
		table: table,

		header:    headerModel{},
		help:      help.New(),
		statusBar: newStatusBarModel(),
		statusLog: newStatusLogModel(),
		syncBar:   newSyncBarModel(),
	}
}

func (m CoreUI) Init() tea.Cmd {
	return tea.SetWindowTitle("Kibe UI")
}

func (m CoreUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.QuitMsg:
		return m, tea.Quit
	case statusLogMessage:
		return m.updateStatusLog(msg, -1), nil
	case statusBarUpdated:
		m.statusBar.SetContent(
			"Resource", msg.resource,
			fmt.Sprintf("Context: %s", msg.context),
			fmt.Sprintf("Namespace: %s", msg.namespace))

		m.statusBar, cmd = m.statusBar.Update(msg)
		return m, cmd
	}

	switch m.viewState {
	case showList:
		return m.updateList(msg)
	case showTable:
		m.keys = m.keys.setEnabled(m.table.fullHelp()...)
		return m.updateTable(msg)
	case showTab:
		switch m.tab.tabViewState {
		case contentSelected:
			m.keys = m.keys.setEnabled(m.tab.fullHelp()...)
		case noContentSelected:
			m.keys = m.keys.setEnabled(m.tab.fullHelpWithContentSelected()...)
		}
		return m.updateTab(msg)
	}
	return m, nil
}

func (m CoreUI) showHelpLines(helpBindingLines ...[]key.Binding) []string {
	var helpLines []string

	helpStyle := lipgloss.NewStyle().MarginBottom(1)

	for _, line := range helpBindingLines {
		helpLines = append(helpLines, helpStyle.Render(
			m.help.ShortHelpView(line)))
	}
	return helpLines
}

func (m CoreUI) composedView() string {
	var helpBindingLines [][]key.Binding
	var dimmMainPaginator bool

	blankSpace := lipgloss.NewStyle().
		Height(windowutil.ComputeHeightPercentage(blankSpaceHeightPercentage)).
		Render("")

	switch m.viewState {
	case showTable:
		dimmMainPaginator = false

		helpBindingLines = [][]key.Binding{
			m.table.firstHelpLineView(),
			m.table.secondHelpLineView(),
		}

	case showTab:
		dimmMainPaginator = true

		switch m.tab.tabViewState {
		case noContentSelected:
			helpBindingLines = [][]key.Binding{
				m.tab.firstHelpLineView(),
				m.tab.secondHelpLineView(),
			}
		case contentSelected:
			helpBindingLines = [][]key.Binding{
				m.tab.firstHelpLineViewWithContentSelected(),
				m.tab.secondHelpLineView(),
			}
		}
	}

	tabPanel := lipgloss.JoinHorizontal(
		lipgloss.Bottom,
		m.tabView(),
		m.statusLogView(),
	)

	helpView := lipgloss.JoinVertical(
		lipgloss.Left,
		m.showHelpLines(helpBindingLines...)...)

	leftUtilityPanel := lipgloss.JoinVertical(
		lipgloss.Left,
		m.table.paginator.view(dimmMainPaginator),
		m.syncBarView(),
	)

	bottomPanel := lipgloss.JoinHorizontal(
		lipgloss.Left,
		leftUtilityPanel,
		lipgloss.NewStyle().Width(3).Render(""),
		helpView,
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.headerView(),
		m.tableView(),
		lipgloss.JoinVertical(lipgloss.Left,
			tabPanel,
			blankSpace,
			bottomPanel,
		),
		m.statusBar.View())
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
