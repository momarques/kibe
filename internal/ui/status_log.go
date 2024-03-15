package ui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
	"github.com/samber/lo"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
)

type statusLogModel struct {
	stream.Stream
}

type statusLogMessage struct {
	duration  time.Duration
	text      string
	timestamp time.Time
}

func newStatusLogModel() statusLogModel {
	stream := koazee.StreamOf([]statusLogMessage{})
	// starts with 10 in order to start with a fixed log string size
	stream = stream.With([]statusLogMessage{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}})
	return statusLogModel{
		Stream: stream,
	}
}

func (m CoreUI) logProcess(text string, duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		return statusLogMessage{
			duration:  duration,
			text:      text,
			timestamp: time.Now(),
		}
	}
}

func (m CoreUI) updateStatusLog(msg tea.Msg) tea.Model {
	switch msg := msg.(type) {
	case statusLogMessage:
		m.statusLog.Stream = m.statusLog.Add(msg)
		if total, _ := m.statusLog.Count(); total > 10 {
			_, m.statusLog.Stream = m.statusLog.Pop()
		}
	}
	return m
}

func (s statusLogModel) String() []string {
	logStream := s.Out().Val().([]statusLogMessage)

	return lo.Map(logStream, func(item statusLogMessage, index int) string {
		var duration string
		var text string = item.text
		var timestamp string

		if item.duration > 0 {
			timestamp = item.timestamp.Format(time.DateTime)
			duration = fmt.Sprintf(" %dms", item.duration.Milliseconds())
		}
		return lipgloss.NewStyle().
			Foreground(uistyles.StatusLogMessages[index]).
			Render(timestamp + " " + text + duration)
	})
}

func (m CoreUI) statusLogView() string {
	return lipgloss.NewStyle().
		MarginTop(11).
		MarginLeft(3).
		Render(strings.Join(m.statusLog.String(), "\n"))
}
