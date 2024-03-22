package kube

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
	"github.com/samber/lo"
)

type ResourceDescription interface {
	TabNames() []string
	TabContent() []string
	SubContent(int) []string
}

// ResourceLabels provides a map of labels from the resource
type ResourceLabels map[string]string

func (rl ResourceLabels) TabContent() string {
	t := table.New()
	t.Rows(
		mapToTableRows(
			mapKeysToYamlKeys(rl))...)
	t.StyleFunc(uistyles.ColorizeTabKeys)
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}

// ResourceAnnotations provides a map of annotations from the resource
type ResourceAnnotations map[string]string

func (ra ResourceAnnotations) TabContent() string {
	t := table.New()
	t.Rows(mapToTableRows(
		mapKeysToYamlKeys(ra))...)
	t.StyleFunc(uistyles.ColorizeTabKeys)
	t.Border(lipgloss.HiddenBorder())
	return t.Render()
}

func mapKeysToYamlKeys(m map[string]string) map[string]string {
	return lo.MapKeys(m, func(value string, key string) string {
		return key + ":"
	})
}
