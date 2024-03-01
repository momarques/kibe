package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/kube"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
	windowutil "github.com/momarques/kibe/internal/ui/window_util"
)

const tabViewShowedHeightPercentage int = 36
const tabViewHiddenHeightPercentage int = 44

type tabModel struct {
	Tabs       []string
	TabContent []string
	activeTab  int
	dimm       bool
}

func newTabUI() tabModel {
	return tabModel{}
}

func (m CoreUI) updateTabUI(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.tabKeys.Back):
			m.state = showTable

			return m, nil

		case key.Matches(msg, m.tabKeys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.tabKeys.NextTab):
			m.tabUI.activeTab = min(m.tabUI.activeTab+1, len(m.tabUI.Tabs)-1)
			return m, nil
		case key.Matches(msg, m.tabKeys.PreviousTab):
			m.tabUI.activeTab = max(m.tabUI.activeTab-1, 0)
			return m, nil
		}
	}
	return m, nil
}

func (m CoreUI) viewTabUI() string {
	if m.tabUI.Tabs == nil {
		return lipgloss.NewStyle().Height(windowutil.
			ComputePercentage(windowHeight, tabViewHiddenHeightPercentage)).Render("")
	}

	switch m.state {
	case showTable:
		m.tabUI.dimm = true
	case showTab:
		m.tabUI.dimm = false
	}

	doc := strings.Builder{}

	var activeStyle, inactiveStyle lipgloss.Style = uistyles.NewTabStyle(m.tabUI.dimm)
	var renderedTabs []string

	for i, t := range m.tabUI.Tabs {
		var style lipgloss.Style

		isFirst, isLast, isActive := m.tabUI.getTabPositions(i)

		if isActive {
			style = activeStyle
		} else {
			style = inactiveStyle
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

	tabs := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	windowStyle := uistyles.NewWindowStyle(m.tabUI.dimm)

	doc.WriteString(tabs)
	doc.WriteString("\n")
	doc.WriteString(
		windowStyle.
			Copy().
			Height(windowutil.
				ComputePercentage(windowHeight, tabViewShowedHeightPercentage)).
			Width(
				(lipgloss.Width(tabs) - windowStyle.GetHorizontalFrameSize())).
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
			pod.Volumes.TabContent(),
			pod.Containers.TabContent(),
			pod.NodeSelectors.TabContent(),
			pod.Tolerations.TabContent(),
			"",
		}
	case *kube.Namespace:
		return nil, nil
	case *kube.Service:
		return nil, nil
	}
	return nil, nil
}
