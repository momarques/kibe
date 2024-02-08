package kubecontext

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/bindings"
	"gopkg.in/yaml.v2"
	"k8s.io/client-go/util/homedir"
)

func NewContextBindings() *contextBindings {
	return &contextBindings{
		chooseContext: bindings.New("enter", "choose"),
	}
}

type contextBindings struct {
	chooseContext key.Binding
}

func (cb *contextBindings) ShortHelpFunc() func() []key.Binding {
	return func() []key.Binding {
		return []key.Binding{cb.chooseContext}
	}
}

func (cb *contextBindings) FullHelpFunc() func() [][]key.Binding {
	return func() [][]key.Binding {
		return [][]key.Binding{cb.ShortHelpFunc()()}
	}
}

func (cb *contextBindings) UpdateFunc() func(msg tea.Msg, m *list.Model) tea.Cmd {
	return func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(contextItem); ok {
			title = i.Title()
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, cb.chooseContext):

				return m.NewStatusMessage(statusMessageStyle("You chose " + title))
			}
		}
		return nil
	}
}

func (cb *contextBindings) FetchListItems() ([]list.Item, error) {
	var kubeconfig string
	var fileContent []byte
	var config = map[string]interface{}{}

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	fileContent, err := os.ReadFile(kubeconfig)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(fileContent, &config); err != nil {
		return nil, err
	}
	return newContextList(config), nil
}

func (cb *contextBindings) NewDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.Styles.SelectedTitle = activeSelectionStyle
	d.Styles.SelectedDesc = activeSelectionStyle

	d.ShortHelpFunc = cb.ShortHelpFunc()
	d.FullHelpFunc = cb.FullHelpFunc()

	d.UpdateFunc = cb.UpdateFunc()

	return d
}
