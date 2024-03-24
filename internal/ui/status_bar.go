package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mistakenelf/teacup/statusbar"
)

type statusBarUpdated struct{ resource, context, namespace string }

func updateStatusBar(r, c, n string) func() tea.Msg {
	return func() tea.Msg {
		return statusBarUpdated{r, c, n}
	}
}

func newStatusBarModel() statusbar.Model {
	s := statusbar.New(
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Light: "#d65f50", Dark: "#d65f50"},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Light: "#3c3836", Dark: "#3c3836"},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Light: "#77352b", Dark: "#77352b"},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Light: "#9b5a46", Dark: "#9b5a46"},
		},
	)
	s.SetContent("Resource", "", "", "")
	return s
}
