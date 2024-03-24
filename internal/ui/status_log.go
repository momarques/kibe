package ui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/logging"
	"github.com/momarques/kibe/internal/ui/style"
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

type loglevel struct {
	level string
	lipgloss.TerminalColor
}

var (
	INFO  = loglevel{"INFO", style.InfoLevel()}
	WARN  = loglevel{"WARN", style.WarnLevel()}
	ERROR = loglevel{"ERROR", style.ErrorLevel()}
	DEBUG = loglevel{"DEBUG", style.DebugLevel()}
)

type statusLogMessage struct {
	operationStatus
	loglevel

	duration  time.Duration
	text      string
	timestamp time.Time
}

func (s statusLogMessage) formatDuration() string {
	if s.duration > 0 {
		return lipgloss.NewStyle().
			Foreground(style.StatusLogDuration()).
			Render(fmt.Sprintf(" duration=%dms", s.duration.Milliseconds()))
	}
	return ""
}

func (s statusLogMessage) formatLogLevel() string {
	return lipgloss.NewStyle().
		Foreground(s.TerminalColor).
		Render(s.level + " ")
}

func (s statusLogMessage) formatTimestamp() string {
	if s.text != "" {
		return fmt.Sprintf("%s ", s.timestamp.Format(time.TimeOnly))
	}
	return ""
}

func (s statusLogMessage) formatStatus() string {
	style := lipgloss.NewStyle().Bold(true)
	switch s.operationStatus {
	case None:
		return ""
	case OK:
		return style.
			Foreground(lipgloss.Color("#a4c847")).
			Render("OK")
	case NOK:
		return style.
			Copy().
			Foreground(lipgloss.Color("#d65f50")).
			Render("NOK")
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
		return item.formatTimestamp() +
			item.formatLogLevel() +
			item.text + " " +
			item.formatStatus() +
			item.formatDuration()
	})
}

func (m CoreUI) logProcess(text string) tea.Cmd {
	return func() tea.Msg {
		return statusLogMessage{
			loglevel:  INFO,
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
