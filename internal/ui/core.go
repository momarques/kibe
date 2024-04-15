package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mistakenelf/teacup/statusbar"
	"github.com/momarques/kibe/internal/kube"
)

const (
	headerSize              = 4
	footerSize              = 6
	tableHeaderSize         = 2
	tableBodySize           = 16
	tableFooterSize         = 1
	tabHeaderSize           = 3
	tabPaginatorSize        = 1
	tabFooterSize           = 1
	tabFooterBlankSpaceSize = 1
)

type viewState int

const (
	showClientConfig viewState = iota
	showTable
	showTab
)

type CoreUI struct {
	viewState

	height int

	client kube.ClientReady

	globalKeys   globalKeyMap
	keys         enabledKeys
	clientConfig clientConfigModel
	tab          tabModel
	table        tableModel

	header    headerModel
	help      help.Model
	log       statusLoggerModel
	statusBar statusbar.Model
	syncBar   syncBarModel
}

func NewUI() CoreUI {
	tab := newTabModel()
	table := newTableModel()

	return CoreUI{
		viewState: showClientConfig,

		client: kube.NewClientReady(),

		globalKeys:   newGlobalKeyMap(),
		keys:         setKeys(table.tableKeyMap, tab.tabKeyMap),
		clientConfig: newClientConfigModel(),
		tab:          tab,
		table:        table,

		header:    headerModel{},
		help:      help.New(),
		log:       newStatusLogger(),
		statusBar: newStatusBarModel(),
		syncBar:   newSyncBarModel(),
	}
}

func (m CoreUI) Init() tea.Cmd {
	return tea.SetWindowTitle("Kibe UI")
}

func (m CoreUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.QuitMsg:
		return m, tea.Quit

	case statusBarUpdated:
		return m.applyStatusBarChanges(msg)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.globalKeys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.globalKeys.SelectContext):
			return m.clearContextSelection(), nil

		case key.Matches(msg, m.globalKeys.SelectNamespace):
			return m.clearNamespaceSelection(), nil

		case key.Matches(msg, m.globalKeys.SelectResource):
			return m.clearResourceSelection(), nil
		}
	}

	switch m.viewState {
	case showClientConfig:
		return m.updateClientConfig(msg)
	case showTable:
		return m.updateTable(msg)
	case showTab:
		return m.updateTab(msg)
	}
	return m, nil
}

func (m CoreUI) showSpecificViewHelpLines(helpBindingLines []key.Binding) string {
	helpStyle := lipgloss.NewStyle().MarginBottom(1)
	return helpStyle.Render(
		m.help.ShortHelpView(helpBindingLines))
}

func (m CoreUI) showGlobalHelpLines() string {
	helpStyle := lipgloss.NewStyle().MarginBottom(1)
	return helpStyle.Render(
		m.help.ShortHelpView(m.globalKeys.fullHelp()))
}

func (m CoreUI) composedView() string {
	var helpBindingLines []key.Binding
	var dimmMainPaginator bool

	blankSpace := lipgloss.NewStyle().
		Height(tabFooterBlankSpaceSize).
		Render("")

	switch m.viewState {
	case showTable:
		dimmMainPaginator = false
		helpBindingLines = m.table.fullHelp()

	case showTab:
		dimmMainPaginator = true

		switch m.tab.tabViewState {
		case noContentSelected:
			helpBindingLines = m.tab.fullHelp()
		case contentSelected:
			helpBindingLines = m.tab.fullHelpWithContentSelected()
		}
	}

	tabPanel := lipgloss.JoinHorizontal(
		lipgloss.Bottom,
		m.tabView(),
		m.statusLogView(),
	)

	helpView := lipgloss.JoinVertical(
		lipgloss.Left,
		m.showSpecificViewHelpLines(helpBindingLines),
		m.showGlobalHelpLines(),
	)

	leftUtilityPanel := lipgloss.JoinVertical(
		lipgloss.Left,
		m.table.paginator.view(dimmMainPaginator),
		lipgloss.NewStyle().Render(""),
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

	case showClientConfig:
		return m.clientConfigView()

	case showTable, showTab:
		return m.composedView()
	}
	return m.View()
}
