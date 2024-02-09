package listmodel

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ListModel struct {
	list list.Model
	keys ListActions
}

func New(listActions ListActions) (ListModel, error) {
	items, err := listActions.FetchListItems()
	if err != nil {
		return ListModel{}, err
	}

	l := list.New(
		items,
		newListOptions(listActions), 0, 0)

	l.Title = listActions.Title()
	l.Styles.Title = titleStyle
	l.Styles.HelpStyle = helpStyle
	l.Styles.FilterPrompt = filterPromptStyle
	l.Styles.FilterCursor = filterCursorStyle

	return ListModel{
		list: l,
		keys: listActions,
	}, nil
}

func (cm ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		cm.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		if cm.list.FilterState() == list.Filtering {
			break
		}
	}
	newListModel, cmd := cm.list.Update(msg)
	cm.list = newListModel
	return cm, cmd
}

func (cm ListModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (cm ListModel) View() string {
	return appStyle.Render(cm.list.View())
}
