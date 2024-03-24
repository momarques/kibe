package style

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
				BorderForeground(GetColor(defaultThemeConfig.Tab.InactiveTabBorder)).
				Padding(0, 1)

	activeTabBorder = tabBorderWithBottom("┘", " ", "└")
	activeTabStyle  = inactiveTabStyle.
			Copy().
			Border(activeTabBorder, true).
			Background(GetColor(defaultThemeConfig.Tab.ActiveTab.BG)).
			Foreground(GetColor(defaultThemeConfig.Tab.ActiveTab.TXT))

	dimmedInactiveTabStyle = lipgloss.NewStyle().
				Border(inactiveTabBorder, true).
				Padding(0, 1).
				BorderForeground(GetColor(defaultThemeConfig.Tab.DimmedInactiveTabBorder)).
				Foreground(GetColor(defaultThemeConfig.Tab.DimmedActiveTab.TXT))

	dimmedActiveTabStyle = dimmedInactiveTabStyle.
				Copy().
				Border(activeTabBorder, true).
				Background(GetColor(defaultThemeConfig.Tab.DimmedActiveTab.BG)).
				Foreground(GetColor(defaultThemeConfig.Tab.DimmedActiveTab.TXT))

	DocStyle = lipgloss.NewStyle().
			Padding(0).
			MarginLeft(2)

	windowStyle = lipgloss.NewStyle().
			BorderForeground(GetColor(defaultThemeConfig.Tab.ActiveTabBorder)).
			Padding(1, 0).
			Border(lipgloss.NormalBorder()).
			UnsetBorderTop()

	dimmedWindowStyle = windowStyle.
				Copy().
				BorderForeground(GetColor(defaultThemeConfig.Tab.DimmedInactiveTabBorder))
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
		return dimmedWindowStyle.Copy()
	}
	return windowStyle.Copy()
}

func NewTabStyle(dimm bool) (lipgloss.Style, lipgloss.Style) {
	if dimm {
		return dimmedActiveTabStyle.Copy(), dimmedInactiveTabStyle.Copy()
	}
	return activeTabStyle.Copy(), inactiveTabStyle.Copy()
}
