package theme

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func Test_measureContentWidthShouldReplaceWithBiggestValue(t *testing.T) {
	expected := "case 3 with bigger string"

	testCases := []string{
		"case 1",
		"case 2",
		expected,
		"case 4 not that big",
		"cas",
	}

	actual := lo.Reduce(testCases, measureContentWidth, 5)
	assert.Equal(t, len(expected), actual)
}

func Test_measureContentWidthShouldKeepInitialValue(t *testing.T) {

	testCases := []string{
		"case 1",
		"case 2",
		"case 3 with bigger string",
		"case 4 not that big",
		"cas",
	}

	expected := 40
	actual := lo.Reduce(testCases, measureContentWidth, 40)
	assert.Equal(t, expected, actual)
}
