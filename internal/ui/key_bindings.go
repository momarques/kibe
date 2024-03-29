package ui

import (
	"slices"

	"github.com/charmbracelet/bubbles/key"
	"github.com/momarques/kibe/internal/bindings"
	"github.com/samber/lo"
)

type globalKeyMap struct {
	SelectContext   key.Binding
	SelectNamespace key.Binding
	SelectResource  key.Binding

	Quit key.Binding
	Help key.Binding
}

func newGlobalKeyMap() globalKeyMap {
	return globalKeyMap{
		SelectContext:   bindings.New("select context", "C", "shift+c"),
		SelectNamespace: bindings.New("select namespace", "N", "shift+n"),
		SelectResource:  bindings.New("select resource", "R", "shift+r"),

		Help: bindings.New("help", "?", "h"),
		Quit: bindings.New("quit", "q", "ctrl+c"),
	}
}

func (k globalKeyMap) fullHelp() []key.Binding {
	return []key.Binding{
		k.SelectContext, k.SelectNamespace, k.SelectResource, k.Help, k.Quit,
	}
}

type tableKeyMap struct {
	Up           key.Binding
	Down         key.Binding
	PreviousPage key.Binding
	NextPage     key.Binding

	CopyID   key.Binding
	ShowYaml key.Binding
	Describe key.Binding
	Delete   key.Binding

	Back key.Binding
	Quit key.Binding
	Help key.Binding
}

func newTableKeyMap() tableKeyMap {
	return tableKeyMap{
		Up:           bindings.New("move up", "up"),
		Down:         bindings.New("move down", "down"),
		PreviousPage: bindings.New("previous page", "left"),
		NextPage:     bindings.New("next page", "right"),
		Describe:     bindings.New("describe resource", "enter", "d"),
	}
}

func (k tableKeyMap) fullHelp() []key.Binding {
	return []key.Binding{
		k.Up, k.Down, k.PreviousPage, k.NextPage, k.Describe,
	}
}

type tabKeyMap struct {
	PreviousTab     key.Binding
	NextTab         key.Binding
	PreviousContent key.Binding
	NextContent     key.Binding

	Choose key.Binding
	Back   key.Binding
}

func newTabKeyMap() tabKeyMap {
	return tabKeyMap{
		PreviousTab:     bindings.New("previous tab", "left", "shift+tab"),
		NextTab:         bindings.New("next tab", "right", "tab"),
		PreviousContent: bindings.New("previous content", "left"),
		NextContent:     bindings.New("next content", "right"),
		Choose:          bindings.New("choose", "enter"),
		Back:            bindings.New("go back", "esc"),
	}
}

func (k tabKeyMap) fullHelp() []key.Binding {
	return []key.Binding{
		k.PreviousTab, k.NextTab, k.Choose, k.Back,
	}
}

func (k tabKeyMap) fullHelpWithContentSelected() []key.Binding {
	return []key.Binding{
		k.PreviousContent, k.NextContent, k.Choose, k.Back,
	}
}

type enabledKeys map[bool][]key.Binding

func setKeys(table tableKeyMap, tab tabKeyMap) enabledKeys {
	var initialKeys []key.Binding

	initialKeys = append(initialKeys, table.fullHelp()...)
	initialKeys = append(initialKeys, tab.fullHelp()...)
	initialKeys = append(initialKeys, tab.fullHelpWithContentSelected()...)

	return enabledKeys{
		false: initialKeys,
	}
}

func (e enabledKeys) setEnabled(keys ...key.Binding) enabledKeys {
	e[true] = lo.Map(keys, func(k key.Binding, _ int) key.Binding {
		k.SetEnabled(true)
		return k
	})

	e[false] = removeKeysFromDisabled(e[false], keys...)
	e[false] = lo.Map(keys, func(k key.Binding, _ int) key.Binding {
		k.SetEnabled(false)
		return k
	})
	return e
}

func removeKeysFromDisabled(disabledKeys []key.Binding, keys ...key.Binding) []key.Binding {
	return lo.DropWhile(keys, func(v key.Binding) bool {
		idx := slices.IndexFunc(disabledKeys, func(k key.Binding) bool {
			return k.Help().Desc == v.Help().Desc
		})
		return idx != -1
	})
}
