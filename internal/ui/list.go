package ui

import (
	"github.com/charmbracelet/bubbles/list"
)

func newListUI(s *selector) list.Model {
	l := list.New(
		[]list.Item{},
		newItemDelegate(s), 0, 0)

	l.Styles.Title = titleStyle
	l.Styles.HelpStyle = helpStyle
	l.Styles.FilterPrompt = filterPromptStyle
	l.Styles.FilterCursor = filterCursorStyle

	return l
}

func (s *selector) setListTitle(m *list.Model, titleMsg string) *list.Model {
	// return func() tea.Msg {

	// }
	m.Title = titleMsg
	return m
}
