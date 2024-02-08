package listmodel

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ContextModel struct {
	list list.Model
	keys ListDelegateBindings
}

func New(listBindings ListDelegateBindings) (ContextModel, error) {
	allContexts, err := listBindings.FetchListItems()
	if err != nil {
		return ContextModel{}, err
	}

	contextList := list.New(
		allContexts,
		listBindings.NewDelegate(), 0, 0)

	contextList.Title = "Choose a context to connect"
	contextList.Styles.Title = titleStyle
	contextList.Styles.HelpStyle = helpStyle
	contextList.Styles.FilterPrompt = filterPromptStyle
	contextList.Styles.FilterCursor = filterCursorStyle

	return ContextModel{
		list: contextList,
		keys: listBindings,
	}, nil
}

func (cm ContextModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (cm ContextModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (cm ContextModel) View() string {
	return appStyle.Render(cm.list.View())
}
