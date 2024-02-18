package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/kube"
	windowutil "github.com/momarques/kibe/internal/ui/window_util"
)

const tabViewProportionPercentage int = 15

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	inactiveTabStyle  = lipgloss.NewStyle().
				Border(inactiveTabBorder, true).
				BorderForeground(highlightColor).Padding(0, 1)

	activeTabBorder = tabBorderWithBottom("┘", " ", "└")
	activeTabStyle  = inactiveTabStyle.
			Copy().
			Border(activeTabBorder, true).
			Background(lipgloss.Color("#ffb1b5")).
			Foreground(lipgloss.Color("#322223"))
	highlightColor = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#ffb8bc"}

	docStyle = lipgloss.NewStyle().
			Padding(1, 2, 1, 2)
	windowStyle = lipgloss.NewStyle().
			BorderForeground(highlightColor).
			Padding(
			windowutil.ComputePercentage(
				windowHeight, tabViewProportionPercentage), 0).
		Align(lipgloss.Center).
		Border(lipgloss.NormalBorder()).
		UnsetBorderTop()
)

type tabModel struct {
	Tabs       []string
	TabContent []string
	activeTab  int
}

func newTabUI() tabModel {
	return tabModel{}
}

// func (t tabModel) colorizeTabContent() string {
// 	colorStyles := uistyles.RandomColorStyleCollection(len(t.TabContent))
// 	var colorizedContent []string

// 	for index, style := range colorStyles {
// 		colorizedContent = append(colorizedContent,
// 			style.Render(t.Tabs[index]),
// 		)
// 	}

// 	return strings.Join(colorizedContent, "\n")

// }

func (m CoreUI) updateTabUI(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "right", "tab":
			m.tabUI.activeTab = min(m.tabUI.activeTab+1, len(m.tabUI.Tabs)-1)
			return m, nil
		case "left", "shift+tab":
			m.tabUI.activeTab = max(m.tabUI.activeTab-1, 0)
			return m, nil
		}

	}
	return m, nil
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

func (m CoreUI) viewTabUI() string {
	doc := strings.Builder{}

	var renderedTabs []string

	for i, t := range m.tabUI.Tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.tabUI.Tabs)-1, i == m.tabUI.activeTab
		if isActive {
			style = activeTabStyle.Copy()
		} else {
			style = inactiveTabStyle.Copy()
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString("\n")
	doc.WriteString(windowStyle.Width((lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize())).Render(m.tabUI.TabContent[m.tabUI.activeTab]))
	return docStyle.Render(doc.String())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (t tabModel) describeResource(c *kube.ClientReady, resourceID string) ([]string, []string) {
	switch c.ResourceSelected.R.(type) {
	case *kube.Pod:
		pod := kube.NewPodDescription(c, resourceID)

		return pod.TabNames(), []string{
			pod.Overview.TabContent(),
			"",
			"",
			"",
			"",
			"",
		}
	case *kube.Namespace:
		return nil, nil
	case *kube.Service:
		return nil, nil
	}
	return nil, nil
}
