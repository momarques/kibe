package listmodel

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/bindings"
	"github.com/momarques/kibe/internal/kube"
)

type Model struct {
	*selector

	List list.Model
}

type selector struct {
	context   string
	namespace string
	resource  string

	choose key.Binding
	client *kube.ClientReady
}

func New(titleMsg string, items []list.Item) (*Model, error) {
	a := selector{
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

	return &Model{
		List:     l,
		selector: &a,
	}, nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		if m.List.FilterState() == list.Filtering {
			break
		}

	case *kube.ClientReady:
		return m, m.selector.clientReady(msg.Resource, msg)

	}
	newListModel, cmd := m.List.Update(msg)
	m.List = newListModel
	return m, cmd
}

func (m Model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m Model) View() string {
	return appStyle.Render(m.List.View())
}
