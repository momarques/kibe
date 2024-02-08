package kubecontext

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type contextModel struct {
	list list.Model
	keys *contextBindings
}

func NewContextModel() (contextModel, error) {
	cb := newContextBindings()
	allContexts, err := fetchAllContexts()
	if err != nil {
		return contextModel{}, err
	}

	contextList := list.New(
		allContexts,
		newContextDelegate(cb), 0, 0)

	contextList.Title = "Choose a context"
	contextList.Styles.Title = titleStyle

	return contextModel{
		list: contextList,
		keys: cb,
	}, nil
}

func (cm contextModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (cm contextModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (cm contextModel) View() string {
	return appStyle.Render(cm.list.View())
}
