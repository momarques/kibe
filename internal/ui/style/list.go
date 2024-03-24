package style

import "github.com/charmbracelet/lipgloss"

var (
	AppStyle = lipgloss.NewStyle().Padding(1, 2)

	ListHeaderTitleStyle = ViewTitleStyle.
				Copy().
				Foreground(GetColor(defaultThemeConfig.List.Header.Title.TXT)).
				Background(GetColor(defaultThemeConfig.List.Header.Title.BG))

	ListActiveSelectionTitleStyle = lipgloss.NewStyle().
					Background(GetColor(defaultThemeConfig.List.ActiveSelectionTitle.BG)).
					Foreground(GetColor(defaultThemeConfig.List.ActiveSelectionTitle.TXT)).
					Padding(0, 1, 0, 1).Bold(true)
	ListActiveSelectionDescStyle = lipgloss.NewStyle().
					Background(GetColor(defaultThemeConfig.List.ActiveSelectionDescription.BG)).
					Foreground(GetColor(defaultThemeConfig.List.ActiveSelectionDescription.TXT)).
					Padding(0, 1, 0, 1)
	ListDimmedDescStyle = lipgloss.NewStyle().
				Foreground(GetColor(defaultThemeConfig.List.DimmedDescription.TXT)).
				Padding(0, 0, 0, 2)
	ListNormalTitleStyle = lipgloss.NewStyle().
				Foreground(GetColor(defaultThemeConfig.List.NormalTitle.TXT)).
				Padding(0, 0, 0, 2)

	ListFilterPromptStyle = lipgloss.NewStyle().
				Foreground(GetColor(defaultThemeConfig.List.Header.FilterPrompt.TXT))

	ListFilterCursorStyle = lipgloss.NewStyle().
				Foreground(GetColor(defaultThemeConfig.List.Header.FilterCursor.TXT))

	StatusMessageStyle = lipgloss.NewStyle().
				Foreground(GetColor(defaultThemeConfig.List.StatusMessage.TXT)).Render
)
