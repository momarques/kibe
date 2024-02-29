package kube

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/momarques/kibe/internal/bindings"
)

type clientKeyMap struct {
	SelectContext   key.Binding
	SelectNamespace key.Binding
	SelectResource  key.Binding
}

var ClientShortcuts = clientKeyMap{
	SelectContext:   bindings.New("select context", "C"),
	SelectNamespace: bindings.New("select namespace", "N"),
	SelectResource:  bindings.New("select resource", "R"),
}
