package style

import (
	"github.com/charmbracelet/lipgloss"
)

func AppStyle() lipgloss.Style {
	return lipgloss.NewStyle().Padding(1, 2)
}

func ListActiveSelectionTitleStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(GetColor(ThemeConfig.List.ActiveSelectionTitle.BG)).
		Foreground(GetColor(ThemeConfig.List.ActiveSelectionTitle.TXT)).
		Padding(0, 1, 0, 1).Bold(true)
}

func ListActiveSelectionDescStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(GetColor(ThemeConfig.List.ActiveSelectionDescription.BG)).
		Foreground(GetColor(ThemeConfig.List.ActiveSelectionDescription.TXT)).
		Padding(0, 1, 0, 1)
}

func ListDimmedDescStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(GetColor(ThemeConfig.List.DimmedDescription.TXT)).
		Padding(0, 0, 0, 2)
}

func ListNormalTitleStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(GetColor(ThemeConfig.List.NormalTitle.TXT)).
		Padding(0, 0, 0, 2)
}

func ListFilterPromptStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(GetColor(ThemeConfig.List.Header.FilterPrompt.TXT))
}

func ListHeaderTitleStyle() lipgloss.Style {
	return ViewTitleStyle.
		Foreground(GetColor(ThemeConfig.List.Header.Title.TXT)).
		Background(GetColor(ThemeConfig.List.Header.Title.BG))
}

func StatusMessageStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(GetColor(ThemeConfig.List.StatusMessage.TXT))
}

func ListFilterCursorStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(GetColor(ThemeConfig.List.Header.FilterCursor.TXT))
}
