package window

import (
	"fmt"
	"testing"
)

func Test_ComputeWidthPercentage(t *testing.T) {
	var testCasesWidth = []int{
		1, 2, 3, 4, 30, 21, 47, 29, 10,
	}

	var windowSize int = 120

	for _, w := range testCasesWidth {
		width := ComputeWidthPercentage(w)
		t.Logf(fmt.Sprintf("%d per cent of %d is equal to %d", w, windowSize, width))
	}
	t.Fail()
}
