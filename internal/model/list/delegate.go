package listmodel

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ListDelegateBindings interface {
	NewDelegate() list.DefaultDelegate

	ShortHelpFunc() func() []key.Binding
	FullHelpFunc() func() [][]key.Binding

	FetchListItems() ([]list.Item, error)
	UpdateFunc() func(msg tea.Msg, m *list.Model) tea.Cmd
}
