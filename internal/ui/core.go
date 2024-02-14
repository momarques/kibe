package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/kube"
)

type coreState int

const (
	showList coreState = iota
	showTable
)

type CoreUI struct {
	state coreState

	client *kube.ClientReady

	listSelector *selector
	listUI       list.Model

	tableContent *content
	tableUI      table.Model
}

func NewUI() CoreUI {
	selector := newListSelector()
	content := newTableContent(nil)

	return CoreUI{
		state:        showList,
		listSelector: selector,
		listUI:       newListUI(selector),
		tableContent: content,
		tableUI:      newTableUI(),
	}
}

func (m CoreUI) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m CoreUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.state {
	case showList:

		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			h, v := appStyle.GetFrameSize()
			m.listUI.SetSize(msg.Width-h, msg.Height-v)

		case tea.KeyMsg:
			if m.listUI.FilterState() == list.Filtering {
				break
			}

		case *kube.ClientReady:
			m.state = showTable
			m.client = msg
			return m, nil
		}

		m.listUI, cmd = m.listUI.Update(msg)
		return m, cmd

	case showTable:

		switch m.tableContent.contentState {
		case loaded:
			switch msg := msg.(type) {
			case tea.KeyMsg:
				if m.listUI.FilterState() == list.Filtering {
					break
				}

				switch msg.String() {
				case "esc":
					if m.tableUI.Focused() {
						m.tableUI.Blur()
					} else {
						m.tableUI.Focus()
					}
				case "q", "ctrl+c":
					return m, tea.Quit
				case "enter":
					return m, tea.Batch(
						tea.Printf("Let's go to %s!", m.tableUI.SelectedRow()[1]),
					)
				}
			}

		case notLoaded:
			m.tableContent.client = m.client
			m.tableUI = *m.tableContent.fetch(&m.tableUI)
		}

		m.tableUI, cmd = m.tableUI.Update(msg)
		return m, cmd
	}
	return nil, nil
}

func (m CoreUI) View() string {
	switch m.state {
	case showList:
		return m.listUI.View()
	case showTable:
		return baseStyle.Render(m.tableUI.View()) + "\n"
	}
	return m.View()
}
