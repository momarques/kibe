package ui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/samber/lo"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
)

type statusLogModel struct {
	logMsgStream stream.Stream
}

type statusLogMessage struct {
	text     string
	duration time.Duration
}

func newStatusLogModel() statusLogModel {
	stream := koazee.StreamOf([]statusLogMessage{})
	// starts with 10 in order to start with a fixed log string size
	stream = stream.With([]statusLogMessage{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}})
	return statusLogModel{
		logMsgStream: stream,
	}
}

func (m CoreUI) logProcess(text string, duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		return statusLogMessage{
			text:     text,
			duration: duration,
		}
	}
}

func (m CoreUI) updateStatusLog(msg tea.Msg) tea.Model {
	switch msg := msg.(type) {
	case statusLogMessage:
		m.logMsgStream = m.logMsgStream.Add(msg)
		if total, _ := m.logMsgStream.Count(); total > 10 {
			_, m.logMsgStream = m.logMsgStream.Pop()
		}
	}
	return m
}

func (s statusLogModel) String() []string {
	logStream := s.logMsgStream.Out().Val().([]statusLogMessage)

	return lo.Map(logStream, func(item statusLogMessage, _ int) string {
		var text = item.text
		var duration string

		if item.duration > 0 {
			duration = fmt.Sprintf(" %dms", item.duration.Milliseconds())
		}
		return text + duration
	})
}

func (m CoreUI) statusLogModelView() string {
	return lipgloss.NewStyle().
		MarginTop(11).
		MarginLeft(3).
		Render(strings.Join(m.statusLogModel.String(), "\n"))
}
