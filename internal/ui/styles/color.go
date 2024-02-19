package uistyles

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/momarques/kibe/internal/logging"
	"github.com/samber/lo"
)

func ColorizeDescriptionSectionKeys(k []string) []string {
	return lo.Map(k, randomColorStyle)
}

func randomColorStyle(content string, _ int) string {
	return resourceSectionDescriptionStyle.
		Copy().
		Background(
			randomColorCode(darkColorCodePrefixes)).
		Foreground(lipgloss.Color("#ffffff")).
		Render(content)
}

var darkColorCodePrefixes = lo.Times(8, func(index int) string {
	return strconv.Itoa(index)
})

const hexCodeChars string = "0123456789abcdef"

func randomColorCode(prefixes []string) lipgloss.Color {
	colorCharacters := lo.Times(5, func(index int) string {
		return string(hexCodeChars[rand.Intn(15)])
	})
	colorCode := "#" + prefixes[rand.Intn(len(prefixes))] +
		strings.Join(colorCharacters, "")
	logging.Log.Info(colorCode)

	return lipgloss.Color(colorCode)
}
