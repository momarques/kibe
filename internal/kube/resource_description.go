package kube

import (
	"github.com/momarques/kibe/internal/ui/style/theme"
	"github.com/samber/lo"
)

type ResourceDescription interface {
	TabNames() []string
	TabContent() []string
	SubContent(int) []string
}

type ResourceLabels map[string]string

func (rl ResourceLabels) TabContent() string {
	keys := lo.Keys(rl)
	content := lo.Values(rl)

	return theme.FormatTable(keys, content)
}

type ResourceAnnotations map[string]string

func (ra ResourceAnnotations) TabContent() string {
	keys := lo.Keys(ra)
	content := lo.Values(ra)

	return theme.FormatTable(keys, content)
}
