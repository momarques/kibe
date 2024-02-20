package uistyles

import (
	"github.com/charmbracelet/lipgloss"
	windowutil "github.com/momarques/kibe/internal/ui/window_util"
)

var (
	WindowWidth, WindowHeight = windowutil.GetWindowSize()
)

const tabViewProportionPercentage int = 10

var (
	resourceSectionDescriptionStyle = lipgloss.NewStyle()

	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	InactiveTabStyle  = lipgloss.NewStyle().
				Border(inactiveTabBorder, true).
				BorderForeground(highlightColor).Padding(0, 1)

	activeTabBorder = tabBorderWithBottom("┘", " ", "└")
	ActiveTabStyle  = InactiveTabStyle.
			Copy().
			Border(activeTabBorder, true).
			Background(lipgloss.Color("#ffb1b5")).
			Foreground(lipgloss.Color("#322223"))
	highlightColor = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#ffb8bc"}

	DocStyle = lipgloss.NewStyle().
			Padding(1, 2, 1, 2)
	WindowStyle = lipgloss.NewStyle().
			BorderForeground(highlightColor).
			Padding(
			windowutil.ComputePercentage(
				WindowHeight, tabViewProportionPercentage), 0).
		Align(lipgloss.Center).
		Border(lipgloss.NormalBorder()).
		UnsetBorderTop()
)

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}
