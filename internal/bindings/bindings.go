package bindings

import "github.com/charmbracelet/bubbles/key"

func New(keyName, actionName string) key.Binding {
	return key.NewBinding(
		key.WithKeys(keyName),
		key.WithHelp(keyName, actionName),
	)
}
