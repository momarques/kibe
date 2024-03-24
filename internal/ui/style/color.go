package style

import (
	"github.com/charmbracelet/lipgloss"
)

func ColorizeTabKeys(row, col int) lipgloss.Style {
	switch {
	case col == 0:
		return resourceSectionDescriptionStyle.
			Copy().
			Background(lipgloss.NoColor{}).
			Foreground(GetColor(defaultThemeConfig.Tab.ActiveTabContentKeys)).
			Padding(0, 1).
			Bold(true)
	}
	return lipgloss.NewStyle()
}

// func colorizeInt(string) string {
// 	return ""
// }

// func colorizeIP(string) string {
// 	return ""
// }

// func colorizeResourceType(string) string {
// 	return ""
// }

// func colorizeStatus(string) string {
// 	return ""
// }

// func colorizePort(string) string {
// 	return ""
// }

// func colorizeBool(string) string {
// 	return ""
// }
