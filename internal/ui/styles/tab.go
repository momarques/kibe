package uistyles

import (
	"github.com/charmbracelet/lipgloss"
	windowutil "github.com/momarques/kibe/internal/ui/window_util"
)

var (
	WindowWidth, WindowHeight = windowutil.GetWindowSize()
)

var (
	resourceSectionDescriptionStyle = lipgloss.NewStyle()

	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	inactiveTabStyle  = lipgloss.NewStyle().
				Border(inactiveTabBorder, true).
				BorderForeground(highlightColor).
				Padding(0, 1)

	activeTabBorder = tabBorderWithBottom("┘", " ", "└")
	activeTabStyle  = inactiveTabStyle.
			Copy().
			Border(activeTabBorder, true).
			Background(lipgloss.Color("#ffb1b5")).
			Foreground(lipgloss.Color("#322223"))

	dimmedInactiveTabStyle = lipgloss.NewStyle().
				Border(inactiveTabBorder, true).
				Padding(0, 1).
				BorderForeground(dimmHighlightColor).
				Foreground(dimmHighlightColor)

	dimmedActiveTabStyle = dimmedInactiveTabStyle.
				Copy().
				Border(activeTabBorder, true).
				Background(dimmHighlightColor).
				Foreground(lipgloss.Color("#aa9890"))

	highlightColor     = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#ffb8bc"}
	dimmHighlightColor = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#4b3e3b"}

	DocStyle = lipgloss.NewStyle().
			Padding(0).
			MarginLeft(2).
			Width(WindowWidth)
	windowStyle = lipgloss.NewStyle().
			BorderForeground(highlightColor).
			Padding(2, 0).
			Align(lipgloss.Center, lipgloss.Center).
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

func NewWindowStyle(dimm bool) lipgloss.Style {
	if dimm {
		return windowStyle.
			Copy().
			BorderForeground(dimmHighlightColor)
	}

	return windowStyle.Copy()
}

func NewTabStyle(dimm bool) (lipgloss.Style, lipgloss.Style) {
	if dimm {
		return dimmedActiveTabStyle.Copy(), dimmedInactiveTabStyle.Copy()
	}
	return activeTabStyle.Copy(), inactiveTabStyle.Copy()
}
