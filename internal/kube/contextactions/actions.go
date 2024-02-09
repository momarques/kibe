package contextactions

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/bindings"
)

func New() *contextActions {
	return &contextActions{
		choose: bindings.New("enter", "choose"),
	}
}

type contextActions struct {
	choose key.Binding
}

func (ca *contextActions) Title() string {
	return "Choose a context to connect"
}

func (ca *contextActions) ShortHelpFunc() func() []key.Binding {
	return func() []key.Binding {
		return []key.Binding{ca.choose}
	}
}

func (ca *contextActions) FullHelpFunc() func() [][]key.Binding {
	return func() [][]key.Binding {
		return [][]key.Binding{ca.ShortHelpFunc()()}
	}
}

func (ca *contextActions) UpdateFunc() func(msg tea.Msg, m *list.Model) tea.Cmd {
	return func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(contextItem); ok {
			title = i.Title()
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, ca.choose):

				return m.NewStatusMessage(statusMessageStyle("You chose " + title))
			}
		}
		return nil
	}
}
