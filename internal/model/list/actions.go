package listmodel

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ListActions interface {
	ShortHelpFunc() func() []key.Binding
	FullHelpFunc() func() [][]key.Binding

	FetchListItems() ([]list.Item, error)
	Title() string
	UpdateFunc() func(msg tea.Msg, m *list.Model) tea.Cmd
}

func newListOptions(l ListActions) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.Styles.SelectedTitle = activeSelectionStyle
	d.Styles.SelectedDesc = activeSelectionStyle

	d.ShortHelpFunc = l.ShortHelpFunc()
	d.FullHelpFunc = l.FullHelpFunc()

	d.UpdateFunc = l.UpdateFunc()

	return d
}
