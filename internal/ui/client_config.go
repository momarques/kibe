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

type clientConfigModel struct {
	list.Model
	clientConfigSelector
}

func newClientConfigModel() clientConfigModel {
	selector := newClientConfigSelector()

	l := list.New(
		[]list.Item{},
		newItemDelegate(selector), 0, 0)

	l.Styles.Title = style.ClientConfigHeaderTitleStyle()
	l.Styles.HelpStyle = style.HelpStyle()
	l.Styles.FilterPrompt = style.ClientConfigFilterPromptStyle()
	l.Styles.FilterCursor = style.ClientConfigFilterCursorStyle()
	l.InfiniteScrolling = false
	l.KeyMap.Quit = bindings.New("quit", "q", "ctrl+c")
	return clientConfigModel{
		Model: l,

		clientConfigSelector: selector,
	}
}

func (m CoreUI) startingTable(c kube.ClientReady) (CoreUI, tea.Cmd) {
	m.viewState = showTable
	m.table.syncState = starting
	m.client = c
	return m, nil
}

func (m CoreUI) updateClientConfig(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := style.
			AppStyle().
			GetFrameSize()
		m.clientConfig.SetSize(msg.Width-h, msg.Height-v)

		m.height = msg.Height
		m.statusBar.SetSize(msg.Width)
		return m, nil

	case tea.KeyMsg:
		if m.clientConfig.FilterState() == list.Filtering {
			break
		}

	case spinner.TickMsg:
		m.clientConfig.spinner, cmd = m.clientConfig.spinner.Update(msg)
		return m, cmd

	case headerTitleUpdated:
		m.header.text = msg
		return m, nil

	case kube.ClientReady:
		m.log.WithDebugContext(m.client).Msg("client is ready")
		return m.startingTable(msg)
	}

	m, cmd = m.clientConfigSelection(msg)
	cmds = append(cmds, cmd)
	m.clientConfig.Model, cmd = m.clientConfig.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m CoreUI) clientConfigView() string {
	if m.clientConfig.spinnerState == showSpinner {
		return lipgloss.JoinVertical(
			lipgloss.Top,
			fmt.Sprintf("%s%s",
				m.clientConfig.spinner.View(),
				m.clientConfig.View()),
			m.statusBar.View())
	}
	return lipgloss.JoinVertical(
		lipgloss.Top, m.clientConfig.View(),
		m.statusBar.View())
}
