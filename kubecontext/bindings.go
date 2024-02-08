package kubecontext

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kb/bindings"
)

func newContextBindings() *contextBindings {
	return &contextBindings{
		chooseContext: bindings.New("enter", "choose"),
	}
}

type contextBindings struct {
	chooseContext key.Binding
}

func (cb *contextBindings) ShortHelpFunc() func() []key.Binding {
	return func() []key.Binding {
		return []key.Binding{cb.chooseContext}
	}
}

func (cb *contextBindings) FullHelpFunc() func() [][]key.Binding {
	return func() [][]key.Binding {
		return [][]key.Binding{cb.ShortHelpFunc()()}
	}
}

func (cb *contextBindings) UpdateFunc() func(msg tea.Msg, m *list.Model) tea.Cmd {
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
			case key.Matches(msg, cb.chooseContext):
				return m.NewStatusMessage(statusMessageStyle("You chose " + title))
			}
		}
		return nil
	}
}

func newContextDelegate(cb *contextBindings) list.DefaultDelegate {
	d := list.NewDefaultDelegate()
	d.UpdateFunc = cb.UpdateFunc()
	d.ShortHelpFunc = cb.ShortHelpFunc()
	d.FullHelpFunc = cb.FullHelpFunc()
	return d
}
