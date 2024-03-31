package ui

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/bindings"
	"github.com/momarques/kibe/internal/kube"
	"github.com/momarques/kibe/internal/logging"
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

	// d.UpdateFunc = s.update

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

func (m CoreUI) clientReady() func() tea.Msg {
	return func() tea.Msg { return m.client.WithContext(context.Background()) }
}

func (m CoreUI) updateClientState() tea.Cmd {
	if m.list.context != "" && m.list.namespace != "" && m.list.resource != "" {
		logging.Log.Info("ready ")
		m.list.clientState = ready
	}

	switch m.list.clientState {
	case ready:
		return m.clientReady()

	case notReady:
		if m.list.context == "" {
			return kube.NewSelectContext()
		}

		if m.list.namespace == "" {
			return kube.NewSelectNamespace(m.client)
		}

		if m.list.resource == "" {
			return kube.NewSelectResource(m.client)
		}
	}
	return nil
}

func (m CoreUI) updateWithKeyStroke(msg tea.KeyMsg) (CoreUI, tea.Cmd) {
	switch {
	case key.Matches(msg, m.list.chooseKey):
		switch i := m.list.SelectedItem().(type) {

		case kube.ContextItem:
			m.list.context = i.FilterValue()
			m.list.spinnerState = showSpinner

			return m, tea.Batch(
				m.list.NewStatusMessage(style.StatusMessageStyle().Render(
					"Context", m.list.context, "set")),
				m.list.contextSelected(m.list.context),
				m.list.spinner.Tick)

		case kube.NamespaceItem:
			m.list.namespace = i.FilterValue()
			m.list.spinnerState = showSpinner

			return m, tea.Batch(
				m.list.NewStatusMessage(style.StatusMessageStyle().Render(
					"Namespace", m.list.namespace, "selected")),
				m.list.namespaceSelected(m.list.namespace),
				m.list.spinner.Tick)

		case kube.ResourceItem:
			m.list.resource = i.FilterValue()
			m.list.spinnerState = showSpinner

			return m, tea.Batch(
				m.list.NewStatusMessage(style.StatusMessageStyle().Render(
					"Showing", m.list.resource+"s")),
				m.list.resourceSelected(),
				m.list.spinner.Tick)
		}
	}
	return m, nil
}

func (m CoreUI) updateListSelector(msg tea.Msg) (CoreUI, tea.Cmd) {
	switch msg := msg.(type) {

	case kube.SelectContext:
		m.list.Title = "Choose a context to connect"
		if m.list.useCurrentContext && msg.CurrentContext != "" && len(msg.Contexts) > 0 {
			m.list.Title = "Skipping context selection"

			m.list.context = msg.CurrentContext
			m.list.spinnerState = showSpinner

			return m, tea.Batch(
				m.list.NewStatusMessage(style.StatusMessageStyle().Render(
					"Using current context", m.list.context)),
				m.list.contextSelected(m.list.context),
				m.list.spinner.Tick)
		}

		if len(msg.Contexts) == 1 {
			m.list.context = msg.Contexts[0].FilterValue()
			m.list.spinnerState = showSpinner

			return m, tea.Batch(
				m.list.NewStatusMessage(style.StatusMessageStyle().Render(
					"Using single existing context", m.list.context)),
				m.list.contextSelected(m.list.context),
				m.list.spinner.Tick)
		}

		m.list.spinnerState = hideSpinner
		return m, m.list.SetItems(msg.Contexts)

	case kube.ContextSelected:
		m.list.ResetFilter()
		m.client = m.client.WithClusterContext(string(msg))

		return m, updateStatusBar(m.list.resource, m.list.context, m.list.namespace)

	case kube.SelectNamespace:
		m.list.Title = "Choose a namespace"
		m.list.spinnerState = hideSpinner

		return m, m.list.SetItems(msg)

	case kube.NamespaceSelected:
		m.list.ResetFilter()
		m.client = m.client.WithNamespace(string(msg))

		return m, updateStatusBar(m.list.resource, m.list.context, m.list.namespace)

	case kube.SelectResource:
		m.list.Title = "Choose a resource type"
		m.list.spinnerState = hideSpinner

		return m, m.list.SetItems(msg.Resources)

	case kube.ResourceSelected:
		m.list.ResetFilter()
		m.client = m.client.WithResource(msg)

		return m, tea.Batch(
			updateStatusBar(m.list.resource, m.list.context, m.list.namespace),
			m.list.updateHeader(fmt.Sprintf("%s interaction", msg.Kind())))

	case tea.KeyMsg:
		return m.updateWithKeyStroke(msg)

	}
	return m, tea.Batch(
		m.updateClientState(),
		m.list.spinner.Tick)
}

func (m CoreUI) cancelTableSync() CoreUI {
	logging.Log.Info("canceling table sync")
	m.client.Cancel()
	m.table.syncState = paused
	return m
}

func (m CoreUI) clearContextSelection() CoreUI {
	m.list.context = ""
	m.list.useCurrentContext = false
	m.list.clientState = notReady
	m.client.ContextSelected = ""
	m.viewState = showList
	return m.clearNamespaceSelection()
}

func (m CoreUI) clearNamespaceSelection() CoreUI {
	m.list.namespace = ""
	m.list.clientState = notReady
	m.client.NamespaceSelected = ""
	m.viewState = showList
	return m.cancelTableSync()
}

func (m CoreUI) clearResourceSelection() CoreUI {
	m.list.resource = ""
	m.list.clientState = notReady
	m.client.ResourceSelected = nil
	m.viewState = showList
	return m.cancelTableSync()
}
