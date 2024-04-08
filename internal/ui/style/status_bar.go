package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mistakenelf/teacup/statusbar"
	"github.com/momarques/kibe/internal/ui/style/theme"
)

func StatusBarFirstColumnColor() statusbar.ColorConfig {
	return statusbar.ColorConfig{
		Foreground: lipgloss.AdaptiveColor{
			Dark:  theme.Selected.StatusBar.ResourceSection.TXT,
			Light: theme.Selected.StatusBar.ResourceSection.TXT,
		},
		Background: lipgloss.AdaptiveColor{
			Light: theme.Selected.StatusBar.ResourceSection.BG,
			Dark:  theme.Selected.StatusBar.ResourceSection.BG,
		},
	}
}

func StatusBarSecondColumnColor() statusbar.ColorConfig {
	return statusbar.ColorConfig{
		Foreground: lipgloss.AdaptiveColor{
			Dark:  theme.Selected.StatusBar.ResourceDetailsSection.TXT,
			Light: theme.Selected.StatusBar.ResourceDetailsSection.TXT,
		},
		Background: lipgloss.AdaptiveColor{
			Light: theme.Selected.StatusBar.ResourceDetailsSection.BG,
			Dark:  theme.Selected.StatusBar.ResourceDetailsSection.BG,
		},
	}
}

func StatusBarThirdColumnColor() statusbar.ColorConfig {
	return statusbar.ColorConfig{
		Foreground: lipgloss.AdaptiveColor{
			Dark:  theme.Selected.StatusBar.ContextSection.TXT,
			Light: theme.Selected.StatusBar.ContextSection.TXT,
		},
		Background: lipgloss.AdaptiveColor{
			Light: theme.Selected.StatusBar.ContextSection.BG,
			Dark:  theme.Selected.StatusBar.ContextSection.BG,
		},
	}
}

func StatusBarFourthColumnColor() statusbar.ColorConfig {
	return statusbar.ColorConfig{
		Foreground: lipgloss.AdaptiveColor{
			Dark:  theme.Selected.StatusBar.NamespaceSection.TXT,
			Light: theme.Selected.StatusBar.NamespaceSection.TXT,
		},
		Background: lipgloss.AdaptiveColor{
			Light: theme.Selected.StatusBar.NamespaceSection.BG,
			Dark:  theme.Selected.StatusBar.NamespaceSection.BG,
		},
	}
}
