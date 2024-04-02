package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/ui/style/theme"
	"github.com/momarques/kibe/internal/ui/style/window"
)

var (
	WindowWidth, WindowHeight = window.GetWindowSize()
	inactiveTabBorder         = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder           = tabBorderWithBottom("┘", " ", "└")
)

func inactiveTabStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(inactiveTabBorder, true).
		BorderForeground(theme.GetColor(theme.Selected.Tab.InactiveTabBorder)).
		Padding(0, 1)
}

func activeTabStyle() lipgloss.Style {
	return inactiveTabStyle().
		Border(activeTabBorder, true).
		Background(theme.GetColor(theme.Selected.Tab.ActiveTab.BG)).
		Foreground(theme.GetColor(theme.Selected.Tab.ActiveTab.TXT))
}

func dimmedInactiveTabStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(inactiveTabBorder, true).
		Padding(0, 1).
		BorderForeground(theme.GetColor(theme.Selected.Tab.DimmedInactiveTabBorder)).
		Foreground(theme.GetColor(theme.Selected.Tab.DimmedActiveTab.TXT))
}

func dimmedActiveTabStyle() lipgloss.Style {
	return dimmedInactiveTabStyle().
		Border(activeTabBorder, true).
		Background(theme.GetColor(theme.Selected.Tab.DimmedActiveTab.BG)).
		Foreground(theme.GetColor(theme.Selected.Tab.DimmedActiveTab.TXT))
}

func DocStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Padding(0).
		MarginLeft(2)
}

func windowStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderForeground(theme.GetColor(theme.Selected.Tab.ActiveTabBorder)).
		Padding(1, 0).
		Border(lipgloss.NormalBorder()).
		UnsetBorderTop()
}

func dimmedWindowStyle() lipgloss.Style {
	return windowStyle().
		BorderForeground(theme.GetColor(theme.Selected.Tab.DimmedInactiveTabBorder))
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

func NewWindowStyle(dimm bool) lipgloss.Style {
	if dimm {
		return dimmedWindowStyle()
	}
	return windowStyle()
}

func NewTabStyle(dimm bool) (lipgloss.Style, lipgloss.Style) {
	if dimm {
		return dimmedActiveTabStyle(), dimmedInactiveTabStyle()
	}
	return activeTabStyle(), inactiveTabStyle()
}
