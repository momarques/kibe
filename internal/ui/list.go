package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mistakenelf/teacup/statusbar"
	"github.com/momarques/kibe/internal/bindings"
	"github.com/momarques/kibe/internal/kube"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

func newListUI(s *selector) list.Model {
	l := list.New(
		[]list.Item{},
		newItemDelegate(s), 0, 0)

	l.Styles.Title = uistyles.ViewTitleStyle.Copy()
	l.Styles.HelpStyle = uistyles.HelpStyle.Copy()
	l.Styles.FilterPrompt = uistyles.ListFilterPromptStyle.Copy()
	l.Styles.FilterCursor = uistyles.ListFilterCursorStyle.Copy()
	l.InfiniteScrolling = false
	l.KeyMap.Quit = bindings.New("q", "quit")

	return l
}

func (m CoreUI) updateListUI(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := uistyles.
			AppStyle.
			Copy().
			GetFrameSize()
		m.listUI.SetSize(msg.Width-h, msg.Height-v)

		m.height = msg.Height
		m.statusbarUI.SetSize(msg.Width)

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

	case UpdateStatusBar:
		var statusbarUI statusbar.Model

		m.statusbarUI.SetContent("Resource",
			m.listSelector.resource,
			fmt.Sprintf("Context: %s", m.listSelector.context),
			fmt.Sprintf("Namespace: %s", m.listSelector.namespace))

		statusbarUI, cmd = m.statusbarUI.Update(msg)
		m.statusbarUI = statusbarUI

		cmds = append(cmds, cmd)
	}

	m.listUI, cmd = m.listUI.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m CoreUI) viewListUI() string {
	if m.listSelector.spinnerState == showSpinner {
		return lipgloss.JoinVertical(
			lipgloss.Top,
			fmt.Sprintf("%s%s",
				m.spinner.View(),
				m.listUI.View()),
			m.statusbarUI.View())
	}
	return lipgloss.JoinVertical(
		lipgloss.Top, m.listUI.View(),
		m.statusbarUI.View())
}
