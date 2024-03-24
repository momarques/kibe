package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/bindings"
	"github.com/momarques/kibe/internal/kube"
	"github.com/momarques/kibe/internal/ui/style"
	"github.com/samber/lo"
)

type clientState int

const (
	ready clientState = iota
	notReady
)

type listSelector struct {
	clientState
	spinnerState

	client    *kube.ClientReady
	spinner   spinner.Model
	chooseKey key.Binding

	context           string
	namespace         string
	resource          string
	useCurrentContext bool
}

func newListSelector() *listSelector {
	sp := spinner.New(
		spinner.WithStyle(style.OKStatusMessage()),
	)
	sp.Spinner = spinner.Dot

	return &listSelector{
		clientState:  notReady,
		spinnerState: hideSpinner,

		useCurrentContext: true,

		chooseKey: bindings.New("choose", "enter"),

		spinner: sp,
	}
}

func newItemDelegate(s *listSelector) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = s.update

	d.Styles.SelectedTitle = style.ListActiveSelectionTitleStyle()
	d.Styles.SelectedDesc = style.ListActiveSelectionDescStyle()
	d.Styles.DimmedDesc = style.ListDimmedDescStyle()
	d.Styles.NormalDesc = style.ListDimmedDescStyle()
	d.Styles.NormalTitle = style.ListNormalTitleStyle()

	d.ShortHelpFunc = func() []key.Binding { return []key.Binding{s.chooseKey} }
	d.FullHelpFunc = func() [][]key.Binding { return [][]key.Binding{{s.chooseKey}} }
	return d
}

func (s listSelector) contextSelected(context string) func() tea.Msg {
	return func() tea.Msg { return kube.ContextSelected(context) }
}

func (s listSelector) namespaceSelected(namespace string) func() tea.Msg {
	return func() tea.Msg { return kube.NamespaceSelected(namespace) }
}

func (s listSelector) resourceSelected() func() tea.Msg {
	r, _ := lo.Find(kube.SupportedResources,
		func(item kube.Resource) bool {
			switch item.Kind() {
			case "Pod", "Service", "Namespace":
				return true
			}
			return false
		})
	return func() tea.Msg { return r }
}

func (s listSelector) clientReady() func() tea.Msg {
	return func() tea.Msg { return s.client }
}

func (s *listSelector) updateClientState() tea.Cmd {
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
	return nil
}

func (s *listSelector) updateWithKeyStroke(msg tea.KeyMsg, m *list.Model) tea.Cmd {
	switch {
	case key.Matches(msg, s.chooseKey):
		switch i := m.SelectedItem().(type) {

		case kube.ContextItem:
			s.context = i.FilterValue()
			s.spinnerState = showSpinner

			return tea.Batch(
				m.NewStatusMessage(style.StatusMessageStyle().Render(
					"Context", s.context, "set")),
				s.contextSelected(s.context),
				s.spinner.Tick)

		case kube.NamespaceItem:
			s.namespace = i.FilterValue()
			s.spinnerState = showSpinner

			return tea.Batch(
				m.NewStatusMessage(style.StatusMessageStyle().Render(
					"Namespace", s.namespace, "selected")),
				s.namespaceSelected(s.namespace),
				s.spinner.Tick)

		case kube.ResourceItem:
			s.resource = i.FilterValue()
			s.spinnerState = showSpinner

			return tea.Batch(
				m.NewStatusMessage(style.StatusMessageStyle().Render(
					"Showing", s.resource+"s")),
				s.resourceSelected(),
				s.spinner.Tick)
		}
	}
	return nil
}

func (s *listSelector) update(msg tea.Msg, m *list.Model) tea.Cmd {
	switch msg := msg.(type) {

	case kube.SelectContext:
		m.Title = "Choose a context to connect"
		if s.useCurrentContext && msg.CurrentContext != "" && len(msg.Contexts) > 0 {
			m.Title = "Skipping context selection"

			s.context = msg.CurrentContext
			s.spinnerState = showSpinner

			return tea.Batch(
				m.NewStatusMessage(style.StatusMessageStyle().Render(
					"Using current context", s.context)),
				s.contextSelected(s.context),
				s.spinner.Tick)
		}

		if len(msg.Contexts) == 1 {
			s.context = msg.Contexts[0].FilterValue()
			s.spinnerState = showSpinner

			return tea.Batch(
				m.NewStatusMessage(style.StatusMessageStyle().Render(
					"Using single existing context", s.context)),
				s.contextSelected(s.context),
				s.spinner.Tick)
		}

		s.spinnerState = hideSpinner
		return m.SetItems(msg.Contexts)

	case kube.ContextSelected:
		m.ResetFilter()

		s.client = kube.NewClientReady(string(msg))
		return updateStatusBar(s.resource, s.context, s.namespace)

	case kube.SelectNamespace:
		m.Title = "Choose a namespace"
		s.spinnerState = hideSpinner

		return m.SetItems(msg)

	case kube.NamespaceSelected:
		m.ResetFilter()
		s.client = s.client.WithNamespace(string(msg))

		return updateStatusBar(s.resource, s.context, s.namespace)

	case kube.SelectResource:
		m.Title = "Choose a resource type"
		s.spinnerState = hideSpinner

		return m.SetItems(msg.Resources)

	case kube.ResourceSelected:
		m.ResetFilter()

		s.client = s.client.WithResource(msg)
		return tea.Batch(
			updateStatusBar(s.resource, s.context, s.namespace),
			s.updateHeader(fmt.Sprintf("%s interaction", msg.Kind())))

	case tea.KeyMsg:
		return s.updateWithKeyStroke(msg, m)

	}
	return tea.Batch(
		s.updateClientState(),
		s.spinner.Tick)
}
