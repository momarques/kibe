package window

import (
	"os"

	"golang.org/x/term"
)

var (
	WindowWidth, WindowHeight = GetWindowSize()
)

func ComputeHeightPercentage(percentage int) int {
	_, h := GetWindowSize()
	return h * percentage / 100
}

func ComputeWidthPercentage(percentage int) int {
	w, _ := GetWindowSize()
	return w * percentage / 100
}

func GetWindowSize() (int, int) {
	isTerminal := term.IsTerminal(int(os.Stdout.Fd()))

	if isTerminal {
		w, h, err := term.GetSize(int(os.Stdout.Fd()))
		if err == nil {
			return w, h
		}
	}
	return 120, 80
}
