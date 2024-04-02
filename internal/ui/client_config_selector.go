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

type clientConfigSelector struct {
	clientState
	spinnerState

	spinner   spinner.Model
	chooseKey key.Binding

	context           string
	namespace         string
	resource          string
	useCurrentContext bool
}

func newClientConfigSelector() clientConfigSelector {
	sp := spinner.New(
		spinner.WithStyle(style.OKStatusMessage()),
	)
	sp.Spinner = spinner.Dot

	return clientConfigSelector{
		clientState:  notReady,
		spinnerState: hideSpinner,

		useCurrentContext: true,

		chooseKey: bindings.New("choose", "enter"),

		spinner: sp,
	}
}

func newItemDelegate(c clientConfigSelector) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.Styles.SelectedTitle = style.ClientConfigActiveSelectionTitleStyle()
	d.Styles.SelectedDesc = style.ClientConfigActiveSelectionDescStyle()
	d.Styles.DimmedDesc = style.ClientConfigDimmedDescStyle()
	d.Styles.NormalDesc = style.ClientConfigDimmedDescStyle()
	d.Styles.NormalTitle = style.ClientConfigNormalTitleStyle()

	d.ShortHelpFunc = func() []key.Binding { return []key.Binding{c.chooseKey} }
	d.FullHelpFunc = func() [][]key.Binding { return [][]key.Binding{{c.chooseKey}} }
	return d
}

func (c clientConfigSelector) contextSelected(context string) func() tea.Msg {
	return func() tea.Msg { return kube.ContextSelected(context) }
}

func (c clientConfigSelector) namespaceSelected(namespace string) func() tea.Msg {
	return func() tea.Msg { return kube.NamespaceSelected(namespace) }
}

func (c clientConfigSelector) resourceSelected() func() tea.Msg {
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
	if m.clientConfig.context != "" && m.clientConfig.namespace != "" && m.clientConfig.resource != "" {
		m.clientConfig.clientState = ready
	}

	switch m.clientConfig.clientState {
	case ready:
		logging.Log.
			WithField("context", m.client.ContextSelected).
			WithField("namespace", m.client.NamespaceSelected).
			WithField("resource", m.client.ResourceSelected).
			Debug("client is ready")
		return m.clientReady()

	case notReady:
		if m.clientConfig.context == "" {
			return kube.NewSelectContext()
		}

		if m.clientConfig.namespace == "" {
			return kube.NewSelectNamespace(m.client)
		}

		if m.clientConfig.resource == "" {
			return kube.NewSelectResource(m.client)
		}
	}
	return nil
}

func (m CoreUI) clientConfigSelected(msg tea.KeyMsg) (CoreUI, tea.Cmd) {
	switch {
	case key.Matches(msg, m.clientConfig.chooseKey):
		switch i := m.clientConfig.SelectedItem().(type) {

		case kube.ContextItem:
			m.clientConfig.context = i.FilterValue()
			m.clientConfig.spinnerState = showSpinner

			return m, tea.Batch(
				m.clientConfig.NewStatusMessage(style.StatusMessageStyle().Render(
					"Context", m.clientConfig.context, "set")),
				m.clientConfig.contextSelected(m.clientConfig.context),
				m.clientConfig.spinner.Tick)

		case kube.NamespaceItem:
			m.clientConfig.namespace = i.FilterValue()
			m.clientConfig.spinnerState = showSpinner

			return m, tea.Batch(
				m.clientConfig.NewStatusMessage(style.StatusMessageStyle().Render(
					"Namespace", m.clientConfig.namespace, "selected")),
				m.clientConfig.namespaceSelected(m.clientConfig.namespace),
				m.clientConfig.spinner.Tick)

		case kube.ResourceItem:
			m.clientConfig.resource = i.FilterValue()
			m.clientConfig.spinnerState = showSpinner

			return m, tea.Batch(
				m.clientConfig.NewStatusMessage(style.StatusMessageStyle().Render(
					"Showing", m.clientConfig.resource+"s")),
				m.clientConfig.resourceSelected(),
				m.clientConfig.spinner.Tick)
		}
	}
	return m, nil
}

func (m CoreUI) clientConfigSelection(msg tea.Msg) (CoreUI, tea.Cmd) {
	switch msg := msg.(type) {

	case kube.SelectContext:
		m.clientConfig.Title = "Choose a context to connect"
		if m.clientConfig.useCurrentContext && msg.CurrentContext != "" && len(msg.Contexts) > 0 {
			m.clientConfig.Title = "Skipping context selection"

			m.clientConfig.context = msg.CurrentContext
			m.clientConfig.spinnerState = showSpinner

			return m, tea.Batch(
				m.clientConfig.NewStatusMessage(style.StatusMessageStyle().Render(
					"Using current context", m.clientConfig.context)),
				m.clientConfig.contextSelected(m.clientConfig.context),
				m.clientConfig.spinner.Tick)
		}

		if len(msg.Contexts) == 1 {
			m.clientConfig.context = msg.Contexts[0].FilterValue()
			m.clientConfig.spinnerState = showSpinner

			return m, tea.Batch(
				m.clientConfig.NewStatusMessage(style.StatusMessageStyle().Render(
					"Using single existing context", m.clientConfig.context)),
				m.clientConfig.contextSelected(m.clientConfig.context),
				m.clientConfig.spinner.Tick)
		}

		m.clientConfig.spinnerState = hideSpinner
		return m, m.clientConfig.SetItems(msg.Contexts)

	case kube.ContextSelected:
		m.clientConfig.ResetFilter()
		m.client = m.client.WithClusterContext(string(msg))

		return m, updateStatusBar(m.clientConfig.resource, m.clientConfig.context, m.clientConfig.namespace)

	case kube.SelectNamespace:
		m.clientConfig.Title = "Choose a namespace"
		m.clientConfig.spinnerState = hideSpinner

		return m, m.clientConfig.SetItems(msg)

	case kube.NamespaceSelected:
		m.clientConfig.ResetFilter()
		m.client = m.client.WithNamespace(string(msg))

		return m, updateStatusBar(m.clientConfig.resource, m.clientConfig.context, m.clientConfig.namespace)

	case kube.SelectResource:
		m.clientConfig.Title = "Choose a resource type"
		m.clientConfig.spinnerState = hideSpinner

		return m, m.clientConfig.SetItems(msg.Resources)

	case kube.ResourceSelected:
		m.clientConfig.ResetFilter()
		m.client = m.client.WithResource(msg)

		return m, tea.Batch(
			updateStatusBar(m.clientConfig.resource, m.clientConfig.context, m.clientConfig.namespace),
			m.clientConfig.updateHeader(fmt.Sprintf("%s interaction", msg.Kind())))

	case tea.KeyMsg:
		return m.clientConfigSelected(msg)

	}
	return m, tea.Batch(
		m.updateClientState(),
		m.clientConfig.spinner.Tick)
}

func (m CoreUI) cancelTableSync() CoreUI {
	logging.Log.
		WithField("context", m.client.ContextSelected).
		WithField("namespace", m.client.NamespaceSelected).
		WithField("resource", m.client.ResourceSelected).
		Debug("canceling table sync")

	m.client.Cancel()
	m.table.syncState = paused
	m.table.response = make(chan kube.TableResponse)
	m.table.columns = nil
	m.table.rows = nil
	return m
}

func (m CoreUI) clearContextSelection() CoreUI {
	m.clientConfig.context = ""
	m.clientConfig.useCurrentContext = false
	m.clientConfig.clientState = notReady
	m.client.ContextSelected = ""
	m.viewState = showClientConfig
	return m.clearNamespaceSelection()
}

func (m CoreUI) clearNamespaceSelection() CoreUI {
	m.clientConfig.namespace = ""
	m.clientConfig.clientState = notReady
	m.client.NamespaceSelected = ""
	m.viewState = showClientConfig
	return m.cancelTableSync()
}

func (m CoreUI) clearResourceSelection() CoreUI {
	m.clientConfig.resource = ""
	m.clientConfig.clientState = notReady
	m.client.ResourceSelected = nil
	m.viewState = showClientConfig
	return m.cancelTableSync()
}
