package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/momarques/kibe/internal/bindings"
)

type tabKeyMap struct {
	PreviousTab key.Binding
	NextTab     key.Binding

	Back key.Binding
	Quit key.Binding
	Help key.Binding
}

func newTabKeyMap() tabKeyMap {
	return tabKeyMap{
		PreviousTab: bindings.New("previous tab", "left", "shift+tab"),
		NextTab:     bindings.New("next tab", "right", "tab"),

		Back: bindings.New("go back", "esc"),
		Help: bindings.New("help", "?", "h"),
		Quit: bindings.New("quit", "q", "ctrl+c"),
	}
}

func (k tabKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Back, k.PreviousTab, k.NextTab, k.Help, k.Quit},
	}
}

func (k tabKeyMap) firstHelpLineView() []key.Binding {
	return []key.Binding{k.Back, k.PreviousTab, k.NextTab, k.Help, k.Quit}
}
