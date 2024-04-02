package theme

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/assert"
)

func Test_GetColor(t *testing.T) {
	testCases := []string{
		"",
		"TESTE",
		"#ffffff",
	}
	expectedColors := []lipgloss.TerminalColor{
		lipgloss.Color("#ffffff"),
		lipgloss.NoColor{},
	}

	for _, color := range testCases {
		got := GetColor(color)
		assert.Contains(t, expectedColors, got)
	}
}
