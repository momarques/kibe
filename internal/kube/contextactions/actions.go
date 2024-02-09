package contextactions

import (
	"context"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/bindings"
	"github.com/momarques/kibe/internal/kube"
	"github.com/momarques/kibe/internal/logging"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func New() *contextActions {
	return &contextActions{
		choose: bindings.New("enter", "choose"),
	}
}

type contextActions struct {
	choose key.Binding
}

func (ca *contextActions) Title() string {
	return "Choose a context to connect"
}

func (ca *contextActions) ShortHelpFunc() func() []key.Binding {
	return func() []key.Binding {
		return []key.Binding{ca.choose}
	}
}

func (ca *contextActions) FullHelpFunc() func() [][]key.Binding {
	return func() [][]key.Binding {
		return [][]key.Binding{ca.ShortHelpFunc()()}
	}
}

func (ca *contextActions) UpdateFunc() func(msg tea.Msg, m *list.Model) tea.Cmd {
	return func(msg tea.Msg, m *list.Model) tea.Cmd {
		var selectedContext string

		if i, ok := m.SelectedItem().(contextItem); ok {
			selectedContext = i.FilterValue()
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, ca.choose):
				client := kube.NewKubeClient(selectedContext)
				ns, err := client.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})
				if err != nil {
					logging.Log.Error(err)
				}
				return m.NewStatusMessage(statusMessageStyle("Namespace " + ns.Items[0].Name))
			}
		}
		return nil
	}
}
