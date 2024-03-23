package ui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/logging"
	uistyles "github.com/momarques/kibe/internal/ui/styles"
	"github.com/samber/lo"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
)

type operationStatus int

const (
	None operationStatus = iota
	NOK
	OK
)

type statusLogMessage struct {
	operationStatus

	duration  time.Duration
	text      string
	timestamp time.Time
}

func (s statusLogMessage) formatDuration() string {
	if s.duration > 0 {
		return fmt.Sprintf(" %dms", s.duration.Milliseconds())
	}
	return ""
}

func (s statusLogMessage) formatTimestamp() string {
	if s.text != "" {
		return fmt.Sprintf("%s ", s.timestamp.Format(time.TimeOnly))
	}
	return ""
}

func (s statusLogMessage) formatStatus() string {
	switch s.operationStatus {
	case None:
		return ""
	case OK:
		return " OK"
	case NOK:
		return " NOK"
	}
	return ""
}

type statusLogModel struct {
	stream.Stream
}

func newStatusLogModel() statusLogModel {
	stream := koazee.StreamOf([]statusLogMessage{})
	// starts with 10 in order to start with a fixed log string size
	stream = stream.With([]statusLogMessage{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}})
	return statusLogModel{
		Stream: stream,
	}
}

func (s statusLogModel) String() []string {
	logStream := s.Out().Val().([]statusLogMessage)

	return lo.Map(logStream, func(item statusLogMessage, index int) string {
		return lipgloss.NewStyle().
			Foreground(uistyles.StatusLogMessages[index]).
			Render(
				item.formatTimestamp() +
					item.text +
					item.formatStatus() +
					item.formatDuration())
	})
}

func (m CoreUI) logProcess(text string) tea.Cmd {
	return func() tea.Msg {
		return statusLogMessage{
			text:      text,
			timestamp: time.Now(),
		}
	}
}

func (m CoreUI) logProcessDuration(status operationStatus, duration time.Duration) (statusLogMessage, int) {
	streamPosition := m.statusLog.Last()
	index, err := m.statusLog.LastIndexOf(streamPosition.Val().(statusLogMessage))
	if err != nil {
		logging.Log.Error(err)
	}

	msg := streamPosition.Val().(statusLogMessage)
	msg.duration = duration
	msg.operationStatus = status

	return msg, index
}

func (m CoreUI) updateStatusLog(msg statusLogMessage, replaceAtIndex int) CoreUI {
	if replaceAtIndex == -1 {
		m.statusLog.Stream = m.statusLog.Add(msg)
		if total, _ := m.statusLog.Count(); total > 10 {
			_, m.statusLog.Stream = m.statusLog.Pop()
		}
	} else {
		m.statusLog.Stream = m.statusLog.Set(replaceAtIndex, msg)
	}
	return m
}

func (m CoreUI) statusLogView() string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		MarginLeft(2).
		Render(strings.Join(m.statusLog.String(), "\n"))
}
