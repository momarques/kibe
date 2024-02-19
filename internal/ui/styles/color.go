package uistyles

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/samber/lo"
)

func ColorizeDescriptionSectionKeys(k []string) []string {
	return lo.Map(k, RandomColorStyle)
}

func RandomColorStyle(content string, _ int) string {
	return resourceSectionDescriptionStyle.
		Copy().
		Background(
			randomColor(darkColorCodePrefixes)).
		Foreground(lipgloss.Color("#ffffff")).
		Render(content)
}

func NewRandomColorStyle() lipgloss.Style {
	return resourceSectionDescriptionStyle.
		Copy().
		Background(
			randomColor(darkColorCodePrefixes)).
		Foreground(lipgloss.Color("#ffffff")).Padding(0, 1)
}

var darkColorCodePrefixes = lo.Times(8, func(index int) string {
	return strconv.Itoa(index)
})

const hexCodeChars string = "0123456789abcdef"

func randomColor(prefixes []string) lipgloss.Color {
	colorCharacters := lo.Times(5, func(index int) string {
		return string(hexCodeChars[rand.Intn(15)])
	})
	colorCode := "#" + prefixes[rand.Intn(len(prefixes))] +
		strings.Join(colorCharacters, "")
	return lipgloss.Color(colorCode)
}

func ColorizeTabKey(row, col int) lipgloss.Style {
	switch {
	case col == 0:
		var foreground string
		foreground = tabKeyPaletteForeground["black"]
		if row > 4 {
			foreground = tabKeyPaletteForeground["light"]
		}

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
