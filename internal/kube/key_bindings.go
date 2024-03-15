package kube

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/momarques/kibe/internal/bindings"
)

type podKeyMap struct {
	ShowNode    key.Binding
	PortForward key.Binding
}

var PodShortcuts = podKeyMap{
	ShowNode:    bindings.New("show node", "n"),
	PortForward: bindings.New("port forward", "p"),
}
