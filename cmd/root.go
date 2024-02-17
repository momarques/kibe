package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/momarques/kibe/internal/logging"
	core "github.com/momarques/kibe/internal/ui"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "kibe",
	Short: "Kubernetes Interaction with Beauty and Elegancy.",
	Long: `
Kibe aims to be a easy tool for interacting with Kubernetes objects without showing a lot of hard to understand information. 
It's a tool focused for developers who doesn't necessarily need to understand the internals of Kubernetes resources.
Also it's a tool made to look beautiful on modern terminals.
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
		const (
			purple    = lipgloss.Color("99")
			gray      = lipgloss.Color("245")
			lightGray = lipgloss.Color("241")
		)

		re := lipgloss.NewRenderer(os.Stdout)

		var (
			HeaderStyle = re.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
			// CellStyle is the base lipgloss style used for the table rows.
			CellStyle = re.NewStyle().Padding(0, 1).Width(14)
			// OddRowStyle is the lipgloss style used for odd-numbered table rows.
			OddRowStyle = CellStyle.Copy().Foreground(gray)
			// EvenRowStyle is the lipgloss style used for even-numbered table rows.
			EvenRowStyle = CellStyle.Copy().Foreground(lightGray)
			// BorderStyle is the lipgloss style used for the table border.
			BorderStyle = lipgloss.NewStyle().Foreground(purple)
		)

		rows := [][]string{
			{"Chinese", "您好", "你好"},
			{"Japanese", "こんにちは", "やあ"},
			{"Arabic", "أهلين", "أهلا"},
			{"Russian", "Здравствуйте", "Привет"},
			{"Spanish", "Hola", "¿Qué tal?"},
		}

		t := table.New().
			Border(lipgloss.NormalBorder()).
			BorderStyle(BorderStyle).
			StyleFunc(func(row, col int) lipgloss.Style {
				var style lipgloss.Style

				switch {
				case row == 0:
					return HeaderStyle
				case row%2 == 0:
					style = EvenRowStyle
				default:
					style = OddRowStyle
				}

				// Make the second column a little wider.
				if col == 1 {
					style = style.Copy().Width(22)
				}

				// Arabic is a right-to-left language, so right align the text.
				if row < len(rows) && rows[row-1][0] == "Arabic" && col != 0 {
					style = style.Copy().Align(lipgloss.Right)
				}

				return style
			}).
			Headers("LANGUAGE", "FORMAL", "INFORMAL").
			Rows(rows...)

		// You can also add tables row-by-row
		t.Row("English", "You look absolutely fabulous.", "How's it going?")

		fmt.Println(t)
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(testCmd)
}
