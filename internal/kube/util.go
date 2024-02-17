package kube

import (
	"strings"
	"time"
)

func DeltaTime(t time.Time) string {
	elapsedTime := time.Since(t)
	elapsedTimeString := elapsedTime.String()

	elapsed, _, _ := strings.Cut(elapsedTimeString, ".")
	return elapsed + "s"
}

func computeWidthPercentage(width, percentage int) int {
	return width * percentage / 100
}
