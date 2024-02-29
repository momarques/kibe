package bindings

import "github.com/charmbracelet/bubbles/key"

func New(actionName string, keyBindings ...string) key.Binding {
	return key.NewBinding(
		key.WithKeys(keyBindings...),
		key.WithHelp(keyBindings[0], actionName),
	)
}
