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

	CopyID   key.Binding
	ShowYaml key.Binding
	Describe key.Binding
	Delete   key.Binding

	Quit key.Binding
	Help key.Binding
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
	Up:           bindings.New("move up", "up"),
	Down:         bindings.New("move down", "down"),
	PreviousPage: bindings.New("previous page", "left"),
	NextPage:     bindings.New("next page", "right"),
	Describe:     bindings.New("describe resource", "enter", "d"),

	Help: bindings.New("help", "?", "h"),
	Quit: bindings.New("quit", "q"),
}

func (k tableKeyMap) viewFirstLine() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.PreviousPage, k.NextPage}
}

func (k tableKeyMap) viewSecondLine() []key.Binding {
	return []key.Binding{k.Describe, k.Help, k.Quit}
}
