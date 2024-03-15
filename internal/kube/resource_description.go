package kube

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

// ResourceLabels provides a map of labels from the resource
type ResourceLabels map[string]string

func (rl ResourceLabels) TabContent() string {
	t := table.New()
	t.Rows(mapToTableRows(rl)...)
	t.StyleFunc(uistyles.ColorizeTabKey)
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}

// ResourceAnnotations provides a map of annotations from the resource
type ResourceAnnotations map[string]string

func (ra ResourceAnnotations) TabContent() string {
	t := table.New()
	t.Rows([]string{})
	t.StyleFunc(uistyles.ColorizeTabKey)
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}
