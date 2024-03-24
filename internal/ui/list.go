package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/bindings"
	"github.com/momarques/kibe/internal/kube"
	"github.com/momarques/kibe/internal/ui/style"
)

type listModel struct {
	list.Model
	*listSelector
}

func newListModel() listModel {
	selector := newListSelector()

	l := list.New(
		[]list.Item{},
		newItemDelegate(selector), 0, 0)

	l.Styles.Title = style.ListHeaderTitleStyle.Copy()
	l.Styles.HelpStyle = style.HelpStyle.Copy()
	l.Styles.FilterPrompt = style.ListFilterPromptStyle.Copy()
	l.Styles.FilterCursor = style.ListFilterCursorStyle.Copy()
	l.InfiniteScrolling = false
	l.KeyMap.Quit = bindings.New("quit", "q", "ctrl+c")
	return listModel{
		Model: l,

		listSelector: selector,
	}
}

func (m CoreUI) updateList(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := style.
			AppStyle.
			Copy().
			GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

		m.height = msg.Height
		m.statusBar.SetSize(msg.Width)
		return m, nil

	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			break
		}

	case spinner.TickMsg:
		m.list.spinner, cmd = m.list.spinner.Update(msg)
		return m, cmd

	case headerTitleUpdated:
		m.header.text = msg
		return m, nil

	case *kube.ClientReady:
		m.viewState = showTable
		m.client = msg
		return m, nil
	}

	m.list.Model, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m CoreUI) listView() string {
	if m.list.spinnerState == showSpinner {
		return lipgloss.JoinVertical(
			lipgloss.Top,
			fmt.Sprintf("%s%s",
				m.list.spinner.View(),
				m.list.View()),
			m.statusBar.View())
	}
	return lipgloss.JoinVertical(
		lipgloss.Top, m.list.View(),
		m.statusBar.View())
}
