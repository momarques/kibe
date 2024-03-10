package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/bindings"
	"github.com/momarques/kibe/internal/kube"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

func newlistModel(s *listSelector) list.Model {
	l := list.New(
		[]list.Item{},
		newItemDelegate(s), 0, 0)

	l.Styles.Title = uistyles.ViewTitleStyle.Copy()
	l.Styles.HelpStyle = uistyles.HelpStyle.Copy()
	l.Styles.FilterPrompt = uistyles.ListFilterPromptStyle.Copy()
	l.Styles.FilterCursor = uistyles.ListFilterCursorStyle.Copy()
	l.InfiniteScrolling = false
	l.KeyMap.Quit = bindings.New("quit", "q", "ctrl+c")
	return l
}

func (m CoreUI) updatelistModel(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := uistyles.
			AppStyle.
			Copy().
			GetFrameSize()
		m.listModel.SetSize(msg.Width-h, msg.Height-v)

		m.height = msg.Height
		m.statusbarModel.SetSize(msg.Width)

	case tea.KeyMsg:
		if m.listModel.FilterState() == list.Filtering {
			break
		}

	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case headerTitleUpdated:
		m.headerModel.text = msg
		return m, nil

	case *kube.ClientReady:
		m.viewState = showTable
		m.client = msg
		return m, nil

	case statusBarUpdated:
		m.statusbarModel.SetContent(
			"Resource", m.listSelector.resource,
			fmt.Sprintf("Context: %s", m.listSelector.context),
			fmt.Sprintf("Namespace: %s", m.listSelector.namespace))

		m.statusbarModel, cmd = m.statusbarModel.Update(msg)
		cmds = append(cmds, cmd)
	}

	m.listModel, cmd = m.listModel.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m CoreUI) listModelView() string {
	if m.listSelector.spinnerState == showSpinner {
		return lipgloss.JoinVertical(
			lipgloss.Top,
			fmt.Sprintf("%s%s",
				m.spinner.View(),
				m.listModel.View()),
			m.statusbarModel.View())
	}
	return lipgloss.JoinVertical(
		lipgloss.Top, m.listModel.View(),
		m.statusbarModel.View())
}
