package windowutil

import (
	"os"

	"golang.org/x/term"
)

func ComputePercentage(target, percentage int) int {
	return target * percentage / 100
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
