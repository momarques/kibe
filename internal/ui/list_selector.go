package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/bindings"
	"github.com/momarques/kibe/internal/kube"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
	"github.com/samber/lo"
)

type clientState int

const (
	ready clientState = iota
	notReady
)

type selector struct {
	clientState

	client *kube.ClientReady

	context   string
	namespace string
	resource  string

	choose key.Binding
}

func newListSelector() *selector {
	return &selector{
		clientState: notReady,
		choose:      bindings.New("enter", "choose"),
	}
}

func newItemDelegate(s *selector) list.DefaultDelegate {
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

	case kube.SelectContext:
		m = s.setListTitle(m, "Choose the context to connect")
		return m.SetItems(msg.Contexts)
	case kube.SelectNamespace:
		m = s.setListTitle(m, "Choose the namespace")
		return m.SetItems(msg.Namespaces)
	case kube.SelectResource:
		m = s.setListTitle(m, "Choose the resource")
		return m.SetItems(msg.Resources)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.choose):
			switch i := m.SelectedItem().(type) {

			case kube.ContextItem:
				s.context = i.FilterValue()

				return tea.Batch(
					m.NewStatusMessage(uistyles.StatusMessageStyle(
						s.context+" selected")),
					s.contextSelected(s.context))

			case kube.NamespaceItem:
				s.namespace = i.FilterValue()

				return tea.Batch(
					m.NewStatusMessage(uistyles.StatusMessageStyle(
						s.namespace+" selected")),
					s.namespaceSelected(s.namespace))

			case kube.ResourceItem:
				s.resource = i.FilterValue()

				return tea.Batch(
					m.NewStatusMessage(uistyles.StatusMessageStyle(
						s.resource+" selected")),
					s.resourceSelected(s.resource))
			}
		}

	case kube.ContextSelected:
		m.ResetFilter()

		s.client = kube.NewClientReady(msg.C)
		return nil

	case kube.NamespaceSelected:
		m.ResetFilter()
		s.client = s.client.WithNamespace(msg.NS)
		return nil

	case kube.ResourceSelected:
		m.ResetFilter()

		s.client = s.client.WithResource(msg.R)
		return nil

	default:

		if s.context != "" && s.namespace != "" && s.resource != "" {
			s.clientState = ready
		}

		switch s.clientState {

		case ready:
			return s.clientReady()

		case notReady:
			if s.context == "" {
				return kube.NewSelectContext()
			}

			if s.namespace == "" {
				return kube.NewSelectNamespace(s.client)
			}

			if s.resource == "" {
				return kube.NewSelectResource(s.client)
			}
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
		return kube.NamespaceSelected{NS: namespace}
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
		return kube.ResourceSelected{R: r}
	}
}

func (s *selector) clientReady() func() tea.Msg {
	return func() tea.Msg {
		return s.client
	}
}
