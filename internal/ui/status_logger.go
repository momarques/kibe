package ui

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/kube"
	"github.com/rs/zerolog"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
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

func (l *logWriter) Write(p []byte) (n int, err error) {
	l.Stream = l.Add(string(p))
	if total, _ := l.Count(); total > 10 {
		_, l.Stream = l.Pop()
	}
	return len(p), nil
}

func formatLogAttrValue(i interface{}) string {
	var text string

	switch value := i.(type) {
	case string:
		text = value
	case int64, int:
		text = fmt.Sprintf("%d", value)
	case float32, float64:
		text = fmt.Sprintf("%1.f", value)
	case json.Number:
		v, _ := value.Float64()
		text = fmt.Sprintf("%1.fms", v)

	case time.Duration:
		text = fmt.Sprintf("%vms", value.Round(time.Millisecond))
	}
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#d8b0f9")).
		Render(fmt.Sprintf("=%v", text))
}

type statusLoggerModel struct {
	zerolog.Logger
	writer *logWriter
}

func newStatusLogger() statusLoggerModel {
	var writers []io.Writer

	if os.Getenv("DEBUG") == "1" {
		file, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		debugWriter := &zerolog.FilteredLevelWriter{
			Writer: zerolog.LevelWriterAdapter{
				Writer: file,
			},
			Level: zerolog.DebugLevel,
		}
		writers = append(writers, debugWriter)
	}

	w := newlogWriter()
	consoleWriter := &zerolog.FilteredLevelWriter{
		Writer: zerolog.LevelWriterAdapter{
			Writer: zerolog.ConsoleWriter{
				Out:        w,
				TimeFormat: time.TimeOnly,
				FormatFieldName: func(i interface{}) string {
					return lipgloss.NewStyle().
						Foreground(lipgloss.Color("#d8b0f9")).
						Render(i.(string))
				},
				FormatFieldValue: formatLogAttrValue,
			},
		},
		Level: zerolog.InfoLevel,
	}
	writers = append(writers, consoleWriter)

	multiWriter := zerolog.MultiLevelWriter(writers...)

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

func (s statusLoggerModel) WithDebugContext(client kube.ClientReady) *zerolog.Event {
	return s.Logger.Debug().
		Caller().
		Dict("client",
			zerolog.Dict().
				Str("context", client.ContextSelected.String()).
				Str("namespace", client.NamespaceSelected.String()).
				Str("resource", client.ResourceSelected.Kind()))
}
