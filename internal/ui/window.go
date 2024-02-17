package ui

import (
	"os"

	"golang.org/x/term"
)

var (
	windowWidth, windowHeight = getWindowSize()
)

func getWindowSize() (int, int) {
	isTerminal := term.IsTerminal(int(os.Stdout.Fd()))

	if isTerminal {
		w, h, err := term.GetSize(int(os.Stdout.Fd()))
		if err == nil {
			return w, h
		}
	}
	return 120, 80
}
