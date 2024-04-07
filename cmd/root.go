package cmd

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/momarques/kibe/internal/logging"
	core "github.com/momarques/kibe/internal/ui"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
)

var RootCmd = &cobra.Command{
	Use:   "kibe",
	Short: "Kubernetes Interaction with Beauty and Elegancy.",
	Long: `
Kibe aims to be an easy and beautiful tool for interacting with Kubernetes objects on modern terminals.
`,
}

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r", "ru"},
	Short:   "Initialize kibe main UI.",
	Run: func(cmd *cobra.Command, args []string) {
		program := tea.NewProgram(
			core.NewUI(),
			tea.WithAltScreen())

		if len(os.Getenv("DEBUG")) > 0 {
			f, err := tea.LogToFile(logging.LogFile, "debug")
			if err != nil {
				fmt.Println("failed to set log file:", err)
				os.Exit(1)
			}
			defer f.Close()
		}

		if _, err := program.Run(); err != nil {
			fmt.Println("failed to run program:", err)
			os.Exit(1)
		}
	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Used for testing layouts without needing to execute the whole program",
	Run: func(cmd *cobra.Command, args []string) {

		// logger := logrus.New()
		// logger.SetOutput(io.Discard)

		// b := bytes.NewBuffer([]byte{})

		// hook := &writer.Hook{
		// 	Writer: b,
		// 	LogLevels: []logrus.Level{
		// 		logrus.PanicLevel,
		// 		logrus.FatalLevel,
		// 		logrus.ErrorLevel,
		// 		logrus.WarnLevel,
		// 	},
		// }
		// logger.SetFormatter(&logrus.TextFormatter{})
		// logger.AddHook(hook)
		// logger.AddHook(&writer.Hook{ // Send info and debug logs to stdout
		// 	Writer: os.Stdout,
		// 	LogLevels: []logrus.Level{
		// 		logrus.InfoLevel,
		// 		logrus.DebugLevel,
		// 	},
		// })
		// logger.Info("This will go to stdout")
		// logger.Warn("This will go to stderr")

		// fmt.Println(b.ReadString('\n'))
		file, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}

		s := newStatusLogWriter()
		multiWriter := zerolog.MultiLevelWriter(
			&zerolog.FilteredLevelWriter{
				Writer: zerolog.LevelWriterAdapter{
					Writer: zerolog.ConsoleWriter{Out: s},
				},
				Level: zerolog.InfoLevel,
			},
			&zerolog.FilteredLevelWriter{
				Writer: zerolog.LevelWriterAdapter{
					Writer: file,
				},
				Level: zerolog.DebugLevel,
			},
		)

		logger := zerolog.New(multiWriter).
			With().
			Timestamp().
			Logger()

		logger.
			Info().
			Msg("hello world")
		logger.Info().Msg("hello world1")
		logger.Info().Msg("hello world2")
		logger.Info().
			Str("status", "OK").
			Str("name", "Tom").
			Msg("hello world3")
		logger.Debug().Msg("hello world4")
		logger.Debug().Msg("hello world5")
		logger.Warn().
			Str("status", "OK").
			Str("name", "Tom").
			Msg("hello world6")
		logger.Warn().Msg("hello world7")
		logger.Info().Msg("hello world8")
		logger.Info().Msg("hello world9")
		logger.Info().Msg("hello world11")

		fmt.Println(strings.Join(s.Out().Val().([]string), ""))

		// l := log.New(s)

		// l.Info("hello", "name", "Al")
		// l.Error("oops", net.ErrClosed, "status", 500)

		// for {
		// 	line, err := s.buf.ReadString('\n')
		// 	if err != nil {
		// 		fmt.Println("failed to read line:", err)
		// 		os.Exit(1)
		// 	}
		// 	fmt.Println(line)

		// }
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(testCmd)
}

type statusLogWriter struct {
	stream.Stream
}

func newStatusLogWriter() *statusLogWriter {
	stream := koazee.StreamOf([]string{})
	return &statusLogWriter{
		Stream: stream.With([]string{"", "", "", "", "", "", "", "", "", ""}),
	}
}

func (s *statusLogWriter) Write(p []byte) (n int, err error) {
	s.Stream = s.Add(string(p))
	if total, _ := s.Count(); total > 10 {
		_, s.Stream = s.Pop()
	}
	return len(p), nil
}

// type Options struct {
// 	// Level reports the minimum level to log.
// 	// Levels with lower levels are discarded.
// 	// If nil, the Handler uses [slog.LevelInfo].
// 	Level slog.Leveler
// }

// type statusHandler struct {
// 	opts Options
// 	// TODO: state for WithGroup and WithAttrs
// 	mu  *sync.Mutex
// 	out io.Writer
// }

// func New(out io.Writer, opts *Options) *statusHandler {
// 	h := &statusHandler{out: out, mu: &sync.Mutex{}}
// 	if opts != nil {
// 		h.opts = *opts
// 	}
// 	if h.opts.Level == nil {
// 		h.opts.Level = slog.LevelInfo
// 	}
// 	return h
// }

// func (s *statusHandler) Enabled(ctx context.Context, level slog.Level) bool {
// 	return true
// }

// func (s *statusHandler) Handle(ctx context.Context, r slog.Record) error {
// 	buf := make([]byte, 0, 1024)
// 	if !r.Time.IsZero() {
// 		buf = s.appendAttr(buf, slog.Time(slog.TimeKey, r.Time), 0)
// 	}
// 	buf = s.appendAttr(buf, slog.Any(slog.LevelKey, r.Level), 0)
// 	if r.PC != 0 {
// 		fs := runtime.CallersFrames([]uintptr{r.PC})
// 		f, _ := fs.Next()
// 		buf = s.appendAttr(buf, slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", f.File, f.Line)), 0)
// 	}
// 	buf = s.appendAttr(buf, slog.String(slog.MessageKey, r.Message), 0)
// 	indentLevel := 0
// 	// TODO: output the Attrs and groups from WithAttrs and WithGroup.
// 	r.Attrs(func(a slog.Attr) bool {
// 		buf = s.appendAttr(buf, a, indentLevel)
// 		return true
// 	})
// 	buf = append(buf, "---\n"...)
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	_, err := s.out.Write(buf)
// 	return err
// }

// func (s statusHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
// 	return nil
// }

// func (s statusHandler) WithGroup(name string) slog.Handler {
// 	return nil
// }
