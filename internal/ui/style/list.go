package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/ui/style/theme"
)

func AppStyle() lipgloss.Style {
	return lipgloss.NewStyle().Padding(1, 2)
}

func ClientConfigActiveSelectionTitleStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(theme.GetColor(theme.Selected.ClientConfig.ActiveSelectionTitle.BG)).
		Foreground(theme.GetColor(theme.Selected.ClientConfig.ActiveSelectionTitle.TXT)).
		Padding(0, 1, 0, 1).Bold(true)
}

func ClientConfigActiveSelectionDescStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(theme.GetColor(theme.Selected.ClientConfig.ActiveSelectionDescription.BG)).
		Foreground(theme.GetColor(theme.Selected.ClientConfig.ActiveSelectionDescription.TXT)).
		Padding(0, 1, 0, 1)
}

func ClientConfigDimmedDescStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.GetColor(theme.Selected.ClientConfig.DimmedDescription.TXT)).
		Padding(0, 0, 0, 2)
}

func ClientConfigNormalTitleStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.GetColor(theme.Selected.ClientConfig.NormalTitle.TXT)).
		Padding(0, 0, 0, 2)
}

func ClientConfigFilterPromptStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.GetColor(theme.Selected.ClientConfig.Header.FilterPrompt.TXT))
}

func ClientConfigHeaderTitleStyle() lipgloss.Style {
	return ViewTitleStyle.
		Foreground(theme.GetColor(theme.Selected.ClientConfig.Header.Title.TXT)).
		Background(theme.GetColor(theme.Selected.ClientConfig.Header.Title.BG))
}

func StatusMessageStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.GetColor(theme.Selected.ClientConfig.StatusMessage.TXT))
}

func ClientConfigFilterCursorStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(theme.GetColor(theme.Selected.ClientConfig.Header.FilterCursor.TXT))
}
