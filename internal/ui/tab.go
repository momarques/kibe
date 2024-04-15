package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/kube"
	"github.com/momarques/kibe/internal/ui/style"
	"github.com/momarques/kibe/internal/ui/style/window"
)

const tabViewHiddenHeightPercentage int = 44
const tabContentHeightPercentage int = 30

type tabViewState int

const (
	contentSelected tabViewState = iota
	noContentSelected
)

type tabModel struct {
	activeTab        int
	activeSubContent int
	dimm             bool
	Tabs             []string
	TabContent       []string
	TabSubContent    []string

	paginator paginatorModel
	kube.ResourceDescription
	tabKeyMap
	tabViewState
}

func newTabModel() tabModel {
	return tabModel{
		activeSubContent: 0,
		tabKeyMap:        newTabKeyMap(),
		tabViewState:     noContentSelected,

		paginator: newPaginatorModel(1),
	}
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

func (m CoreUI) updateTab(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.tab.tabViewState {
	case noContentSelected:
		m.keys = m.keys.setEnabled(m.tab.fullHelpWithContentSelected()...)

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, m.tab.Back):
				m.viewState = showTable
				m, cmd = m.syncTable()
				return m, tea.Batch(cmd,
					updateStatusBar(m.client.Kind(),
						m.client.ContextSelected.String(),
						m.client.NamespaceSelected.String()))

			case key.Matches(msg, m.tab.NextTab):
				m.tab.activeTab = min(m.tab.activeTab+1, len(m.tab.Tabs)-1)
				return m, nil

			case key.Matches(msg, m.tab.PreviousTab):
				m.tab.activeTab = max(m.tab.activeTab-1, 0)
				return m, nil

			case key.Matches(msg, m.tab.Choose):
				m.tab = m.tab.fetchSubContent(msg)
				return m, nil
			}
		}
		return m, nil

	case contentSelected:
		m.keys = m.keys.setEnabled(m.tab.fullHelp()...)

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, m.tab.Back):
				m.tab.tabViewState = noContentSelected
				m.tab.paginator.Page = 0
				m.tab.activeSubContent = 0
				return m, nil

			case key.Matches(msg, m.tab.NextContent):
				m.tab.paginator.Model, _ = m.tab.paginator.Update(msg)
				m.tab.activeSubContent = min(m.tab.activeSubContent+1, len(m.tab.TabSubContent)-1)
				return m, nil

			case key.Matches(msg, m.tab.PreviousContent):
				m.tab.paginator.Model, _ = m.tab.paginator.Update(msg)
				m.tab.activeSubContent = max(m.tab.activeSubContent-1, 0)
				return m, nil

			case key.Matches(msg, m.tab.Choose):

			}
		}
		return m, nil
	}
	return m, nil
}

func (t tabModel) getTabPositions(index int) (bool, bool, bool) {
	return index == 0, index == len(t.Tabs)-1, index == t.activeTab
}

func (m CoreUI) formatTabs() []string {
	var activeStyle, inactiveStyle lipgloss.Style = style.NewTabStyle(m.tab.dimm)
	var renderedTabs []string

	for i, t := range m.tab.Tabs {
		var style lipgloss.Style

		isFirst, isLast, isActive := m.tab.getTabPositions(i)

		if isActive {
			style = activeStyle.Copy()
		} else {
			style = inactiveStyle.Copy()
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

	return renderedTabs
}

func (m CoreUI) tabView() string {
	_, h := window.GetWindowSize()

	if m.tab.Tabs == nil {
		m.log.Debug().
			Int("height", h).
			Int("tab hidden size", h-showedTabSize()).
			Msg("tabView")
		return lipgloss.NewStyle().
			Height(h - showedTabSize()).
			Width(103).
			Render("")
	}

	switch m.viewState {
	case showTable:
		m.tab.dimm = true
	case showTab:
		m.tab.dimm = false
	}

	tabWindow := strings.Builder{}

	tabs := lipgloss.JoinHorizontal(lipgloss.Top, m.formatTabs()...)
	windowStyle := style.NewWindowStyle(m.tab.dimm)

	contentStyle := windowStyle.
		Copy().
		Width((lipgloss.Width(tabs) - windowStyle.GetHorizontalFrameSize()))

	var content string
	var contentBlockStyle lipgloss.Style = lipgloss.NewStyle().
		Height(h - hiddenTabSize())
		// Width(100)
	var paginatorView string = "\n"

	switch m.tab.tabViewState {
	case noContentSelected:
		content = m.tab.TabContent[m.tab.activeTab]
	case contentSelected:
		content = m.tab.TabSubContent[m.tab.activeSubContent]
		paginatorView = m.tab.paginator.view(m.tab.dimm)
	}

	tabWindow.WriteString(tabs)
	tabWindow.WriteString("\n")
	tabWindow.WriteString(contentStyle.
		Render(lipgloss.JoinVertical(
			lipgloss.Left,
			contentBlockStyle.Render(content),
			paginatorView,
		)))

	m.log.Debug().
		Int("height", h).
		Int("tab showed size", h-hiddenTabSize()).
		Msg("tabView")
	return style.TabWindowStyle().
		Render(tabWindow.String())
}

type descriptionReady struct {
	tabNames   []string
	tabContent []string
}

func (t tabModel) describeResource(c kube.ClientReady) (tabModel, tea.Cmd) {
	t.ResourceDescription = c.ResourceSelected.Describe(c)
	return t, func() tea.Msg {
		return descriptionReady{
			t.ResourceDescription.TabNames(),
			t.ResourceDescription.TabContent(),
		}
	}
}

func (t tabModel) fetchSubContent(msg tea.Msg) tabModel {
	t.TabSubContent = t.ResourceDescription.SubContent(t.activeTab)
	if len(t.TabSubContent) > 0 {
		t.tabViewState = contentSelected
		t.paginator.SetTotalPages(len(t.TabSubContent))
		t.paginator.Model, _ = t.paginator.Update(msg)
	}
	return t
}
