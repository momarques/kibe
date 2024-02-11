package listmodel

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/bindings"
)

type model struct {
	list    list.Model
	actions actions
}

type actions struct {
	selectedContext  string
	selectedResource string
	choose           key.Binding
}

func New(titleMsg string, items []list.Item) (model, error) {
	a := actions{
		choose: bindings.New("enter", "choose"),
	}

	l := list.New(
		items,
		a.newDelegate(), 0, 0)

	l.Title = titleMsg
	l.Styles.Title = titleStyle
	l.Styles.HelpStyle = helpStyle
	l.Styles.FilterPrompt = filterPromptStyle
	l.Styles.FilterCursor = filterCursorStyle

	return model{
		list:    l,
		actions: a,
	}, nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			break
		}

	}

	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	return m, cmd
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) View() string {
	return appStyle.Render(m.list.View())
}
