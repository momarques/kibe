package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/kube"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
)

type tabModel struct {
	Tabs       []string
	TabContent []string
	activeTab  int
}

func newTabUI() tabModel {
	return tabModel{}
}

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

func (m CoreUI) viewTabUI() string {
	doc := strings.Builder{}

	var renderedTabs []string

	for i, t := range m.tabUI.Tabs {
		var style lipgloss.Style

		isFirst, isLast, isActive := m.tabUI.getTabPositions(i)

		if isActive {
			style = uistyles.ActiveTabStyle.Copy()
		} else {
			style = uistyles.InactiveTabStyle.Copy()
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
	doc.WriteString(
		uistyles.WindowStyle.
			Width(
				(lipgloss.Width(row) - uistyles.WindowStyle.GetHorizontalFrameSize())).
			Render(
				m.tabUI.TabContent[m.tabUI.activeTab]))

	return uistyles.DocStyle.Render(doc.String())
}

func (t tabModel) getTabPositions(index int) (bool, bool, bool) {
	return index == 0, index == len(t.Tabs)-1, index == t.activeTab
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
			pod.Status.TabContent(),
			pod.Labels.TabContent(),
			pod.Annotations.TabContent(),
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
