package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/kube"
	"github.com/momarques/kibe/internal/logging"
)

type coreState int

const (
	showList coreState = iota
	showTable
)

type CoreUI struct {
	state coreState

	listSelector *selector
	client       *kube.ClientReady

	listUI  list.Model
	tableUI table.Model
}

func NewUI() CoreUI {
	selector := newListSelector()
	return CoreUI{
		state:        showList,
		listSelector: selector,
		listUI:       newListUI(selector),
		tableUI:      newTableUI(),
	}
}

func (m CoreUI) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m CoreUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	logging.Log.Info("core model")
	var cmd tea.Cmd

	switch m.state {
	case showList:
		m.listUI.Title = ""

		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			h, v := appStyle.GetFrameSize()
			m.listUI.SetSize(msg.Width-h, msg.Height-v)

		case tea.KeyMsg:
			if m.listUI.FilterState() == list.Filtering {
				break
			}

		case *kube.ClientReady:
			logging.Log.Info("passou ->>")
			m.state = showTable
			return m, nil
		}

		m.listUI, cmd = m.listUI.Update(msg)
		return m, cmd

	case showTable:

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

		m.tableUI, cmd = m.tableUI.Update(msg)
		return m, cmd
	}

	return nil, nil
}

func (m CoreUI) View() string {
	logging.Log.Info("estado atual -> ", m.state)

	switch m.state {
	case showList:
		return m.listUI.View()
	case showTable:
		return baseStyle.Render(m.tableUI.View()) + "\n"
	}
	return m.View()
}
