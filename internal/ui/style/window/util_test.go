package window

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ComputeWidthPercentage(t *testing.T) {
	testCasesWidth := []int{
		1, 2, 3, 4, 30, 21, 47, 29, 10,
	}

	expectedComputedValues := []int{
		1, 2, 3, 4, 36, 25, 56, 34, 12,
	}

	for i, w := range testCasesWidth {
		actual := ComputeWidthPercentage(w)
		if !assert.Equal(t, expectedComputedValues[i], actual) {
			t.Logf("Test case: %d -> expected computed value: %d, got %d", i, expectedComputedValues[i], actual)
		}
	}
}

func Test_ComputeHeightPercentage(t *testing.T) {
	testCasesHeight := []int{
		1, 2, 3, 4, 30, 21, 47, 29, 10,
	}

	expectedComputedValues := []int{
		0, 1, 2, 3, 24, 16, 37, 23, 8,
	}

	for i, h := range testCasesHeight {
		actual := ComputeHeightPercentage(h)
		if !assert.Equal(t, expectedComputedValues[i], actual) {
			t.Logf("Test case: %d -> expected computed value: %d, got %d", i, expectedComputedValues[i], actual)
		}
	}
}
