package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/kube"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

type coreState int

const (
	showList coreState = iota
	showTable
)

type spinnerState int

const (
	showSpinner spinnerState = iota
	hideSpinner
)

type CoreUI struct {
	state coreState

	client *kube.ClientReady

	listSelector *selector
	listUI       list.Model

	tableContent *content
	tableUI      table.Model

	spinner spinner.Model
}

func NewUI() CoreUI {
	s := spinner.New(
		spinner.WithStyle(uistyles.OKStatusMessage),
	)
	s.Spinner = spinner.Dot
	selector := newListSelector(s)
	content := newTableContent(nil)

	list := newListUI(selector)
	return CoreUI{
		state:        showList,
		listSelector: selector,
		listUI:       list,
		tableContent: content,
		tableUI:      newTableUI(),
		spinner:      s,
	}
}

func (m CoreUI) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m CoreUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
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

		case spinner.TickMsg:
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd

		case *kube.ClientReady:
			m.state = showTable
			m.client = msg
			return m, nil
		}

		m.listUI, cmd = m.listUI.Update(msg)
		cmds = append(cmds, cmd)

		return m, tea.Batch(cmds...)

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
		if m.listSelector.spinnerState == showSpinner {
			return fmt.Sprintf("%s%s", m.spinner.View(), m.listUI.View())
		}
		return m.listUI.View()
	case showTable:
		return baseStyle.Render(m.tableUI.View()) + "\n"
	}
	return m.View()
}
