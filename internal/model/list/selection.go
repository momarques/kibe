package listmodel

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/kube"
	"github.com/momarques/kibe/internal/logging"
	modelstyles "github.com/momarques/kibe/internal/model/styles"
	"github.com/samber/lo"
)

func (s *selector) newDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = s.update

	d.Styles.SelectedTitle = activeSelectionStyle
	d.Styles.SelectedDesc = activeSelectionStyle

	d.ShortHelpFunc = func() []key.Binding {
		return []key.Binding{s.choose}
	}
	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{{s.choose}}
	}
	return d
}

func (s *selector) update(msg tea.Msg, m *list.Model) tea.Cmd {
	switch msg := msg.(type) {
	case kube.ContextSelected:
		m.ResetFilter()

		s.client = kube.NewClientReady(msg.C)

		items, err := kube.NamespacesAsList(s.client)
		if err != nil {
			logging.Log.Error(err)
		}

		return m.SetItems(items)

	case kube.NamespaceSelected:
		m.ResetFilter()

		items, err := kube.ListAvailableResources(s.client)
		if err != nil {
			logging.Log.Error(err)
		}

		return m.SetItems(items)

	case kube.ResourceSelected:
		m.ResetFilter()

		return s.clientReady(msg.Resource, s.client)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.choose):
			switch i := m.SelectedItem().(type) {
			case kube.SelectContext:
				s.context = i.FilterValue()

				return tea.Batch(
					m.NewStatusMessage(modelstyles.StatusMessageStyle(
						s.context+" selected")),
					s.contextSelected(s.context))

			case kube.SelectNamespace:
				s.namespace = i.FilterValue()

				return tea.Batch(
					m.NewStatusMessage(modelstyles.StatusMessageStyle(
						s.namespace+" selected")),
					s.namespaceSelected(s.namespace))

			case kube.SelectResource:
				s.resource = i.FilterValue()

				return tea.Batch(
					m.NewStatusMessage(modelstyles.StatusMessageStyle(
						s.resource+" selected")),
					s.resourceSelected(s.resource))

			}
			return nil
		}
	}
	return nil
}

func (s *selector) contextSelected(context string) func() tea.Msg {
	return func() tea.Msg {
		return kube.ContextSelected{C: context}
	}
}

func (s *selector) namespaceSelected(namespace string) func() tea.Msg {
	return func() tea.Msg {
		return kube.NamespaceSelected(namespace)
	}
}

func (s *selector) resourceSelected(kind string) func() tea.Msg {
	r, _ := lo.Find(kube.SupportedResources, func(item kube.Resource) bool {
		switch item.Kind() {
		case "Pod", "Service", "Namespace":
			return true
		}
		return false
	})

	return func() tea.Msg {
		return r
	}
}

func (s *selector) clientReady(resource kube.Resource, c *kube.ClientReady) func() tea.Msg {
	return func() tea.Msg {
		return c.SetNamespace(s.namespace).
			SetResource(resource)
	}
}
