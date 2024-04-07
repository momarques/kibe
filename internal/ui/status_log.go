package ui

import (
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/rs/zerolog"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
)

type operationStatus int

const (
	None operationStatus = iota
	NOK
	OK
)

type logWriter struct {
	stream.Stream
}

func newlogWriter() *logWriter {
	stream := koazee.StreamOf([]string{})
	return &logWriter{
		Stream: stream.With([]string{"", "", "", "", "", "", "", "", "", ""}),
	}
}

func (s *logWriter) Write(p []byte) (n int, err error) {
	s.Stream = s.Add(string(p))
	if total, _ := s.Count(); total > 10 {
		_, s.Stream = s.Pop()
	}
	return len(p), nil
}

type statusLoggerModel struct {
	zerolog.Logger
	writer *logWriter
}

func newStatusLogger() statusLoggerModel {
	var loggers []io.Writer

	if os.Getenv("DEBUG") == "1" {
		file, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		debugLogger := &zerolog.FilteredLevelWriter{
			Writer: zerolog.LevelWriterAdapter{
				Writer: file,
			},
			Level: zerolog.DebugLevel,
		}
		loggers = append(loggers, debugLogger)
	}

	w := newlogWriter()
	consoleLogger := &zerolog.FilteredLevelWriter{
		Writer: zerolog.LevelWriterAdapter{
			Writer: zerolog.ConsoleWriter{Out: w},
		},
		Level: zerolog.InfoLevel,
	}
	loggers = append(loggers, consoleLogger)

	multiWriter := zerolog.MultiLevelWriter(loggers...)

	return statusLoggerModel{
		writer: w,
		Logger: zerolog.New(multiWriter).
			With().
			Timestamp().
			Logger(),
	}
}

func (s statusLoggerModel) String() string {
	logStream := s.writer.Out().Val().([]string)
	return strings.Join(logStream, "")
}

func (m CoreUI) statusLogView() string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		MarginLeft(2).
		Render(m.log.String())
}
