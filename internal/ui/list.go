package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/momarques/kibe/internal/bindings"
)

func newListUI(s *selector) list.Model {
	l := list.New(
		[]list.Item{},
		newItemDelegate(s), 0, 0)

	l.Styles.Title = titleStyle
	l.Styles.HelpStyle = helpStyle
	l.Styles.FilterPrompt = filterPromptStyle
	l.Styles.FilterCursor = filterCursorStyle
	l.InfiniteScrolling = false
	l.KeyMap.Quit = bindings.New("q", "quit")

	return l
}
