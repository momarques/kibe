package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/momarques/kibe/internal/bindings"
)

type tabKeyMap struct {
	PreviousTab key.Binding
	NextTab     key.Binding

	Quit key.Binding
	Help key.Binding
}

func (k tabKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.PreviousTab, k.NextTab}
}

func (k tabKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.PreviousTab, k.NextTab, k.Help, k.Quit},
	}
}

var tabShortcuts = tabKeyMap{
	PreviousTab: bindings.New("previous tab", "left", "shift+tab"),
	NextTab:     bindings.New("next tab", "right", "tab"),

	Help: bindings.New("help", "?", "h"),
	Quit: bindings.New("quit", "q", "ctrl+c"),
}

func (k tabKeyMap) viewFirstLine() []key.Binding {
	return []key.Binding{k.PreviousTab, k.NextTab, k.Help, k.Quit}
}
