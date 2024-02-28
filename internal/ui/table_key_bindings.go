package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/momarques/kibe/internal/bindings"
)

type tableKeyMap struct {
	Up           key.Binding
	Down         key.Binding
	PreviousPage key.Binding
	NextPage     key.Binding
	Help         key.Binding
	Quit         key.Binding
	Describe     key.Binding
}

func (k tableKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Up, k.Down, k.PreviousPage, k.NextPage, k.Describe, k.Help, k.Quit}
}

func (k tableKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.PreviousPage, k.NextPage},
		{k.Describe, k.Help, k.Quit},
	}
}

var tableShortcuts = tableKeyMap{
	Up:           bindings.New("up", "move up"),
	Down:         bindings.New("down", "move down"),
	PreviousPage: bindings.New("left", "previous page"),
	NextPage:     bindings.New("right", "next page"),
	Describe:     bindings.New("enter", "describe resource"),
	Help:         bindings.New("?", "help"),
	Quit:         bindings.New("q", "quit"),
}

func (k tableKeyMap) viewFirstLine() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.PreviousPage, k.NextPage}
}

func (k tableKeyMap) viewSecondLine() []key.Binding {
	return []key.Binding{k.Describe, k.Help, k.Quit}
}
