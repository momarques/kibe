package coremodel

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/kube"
	listmodel "github.com/momarques/kibe/internal/model/list"
	tablemodel "github.com/momarques/kibe/internal/model/table"
)

type state int

const (
	showList state = iota
	showTable
)

type model struct {
	state     state
	listView  *listmodel.Model
	tableView tablemodel.Model
}

func New() model {
	contexts, err := kube.ListContexts()
	if err != nil {
		fmt.Printf("failed to create model: %s", err)
		os.Exit(1)
	}

	listView, err := listmodel.New("Choose a context to connect", contexts)
	if err != nil {
		fmt.Printf("failed to create model: %s", err)
		os.Exit(1)
	}

	return model{
		state:    showList,
		listView: listView,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case showList:
		return m.listView.Update(msg)
	case showTable:
		return m.tableView.Update(msg)
	}
	return nil, nil
}

func (m model) View() string {
	switch m.state {
	case showList:
		return m.listView.View()
	case showTable:
		return m.tableView.View()
	}
	return m.View()
}
