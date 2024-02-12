package listmodel

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/kube"
	"github.com/momarques/kibe/internal/logging"
	modelstyles "github.com/momarques/kibe/internal/model/styles"
	"k8s.io/client-go/kubernetes"
)

func (a actions) newDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = a.updateFunc

	d.Styles.SelectedTitle = activeSelectionStyle
	d.Styles.SelectedDesc = activeSelectionStyle

	d.ShortHelpFunc = func() []key.Binding {
		return []key.Binding{a.choose}
	}
	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{{a.choose}}
	}
	return d
}

func (a actions) updateFunc(msg tea.Msg, m *list.Model) tea.Cmd {
	var client *kubernetes.Clientset

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, a.choose):
			switch s := m.SelectedItem().(type) {
			case kube.ContextItem:
				a.selectedContext = s.FilterValue()

				client = kube.NewKubeClient(a.selectedContext)
				apiList, err := client.ServerPreferredResources()
				if err != nil {
					logging.Log.Error(err)
				}
				items, err := kube.FetchListItems(apiList)
				if err != nil {
					logging.Log.Error(err)
				}

				return tea.Batch(
					m.NewStatusMessage(modelstyles.StatusMessageStyle(
						a.selectedContext+" selected")),
					m.SetItems(items))

			case kube.ResourceItem:
				a.selectedResource = s.FilterValue()

				return m.NewStatusMessage(modelstyles.StatusMessageStyle(
					a.selectedResource + " selected"))
			}
			return nil
		}
	}
	return nil
}
