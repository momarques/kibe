package uistyles

import (
	"github.com/charmbracelet/lipgloss"
)

func ColorizeTabKey(row, col int) lipgloss.Style {
	switch {
	case col == 0:
		var foreground string
		foreground = tabKeyPaletteForeground["black"]
		// if row > 4 {
		// 	foreground = tabKeyPaletteForeground["light"]
		// }

		return resourceSectionDescriptionStyle.
			Copy().
			Background(lipgloss.Color(tabKeyPaletteBackground[row])).
			Foreground(lipgloss.Color(foreground)).
			Padding(0, 1).
			Bold(true)
	}
	return lipgloss.NewStyle()
}

var tabKeyPaletteForeground = map[string]string{
	"light": "#ffffff",
	"dark":  "#210f0d",
}

var tabKeyPaletteBackground = []string{
	// light colors
	"#ffa598",
	"#ff9b8e",
	"#ff9184",
	"#ff877a",
	// dark colors
	"#ff7d70",
	"#ff7366",
	"#f5695c",
	"#eb5f52",
	"#e15548",
	"#d74b3e",
	"#cd4134",
	"#c3372a",
}
