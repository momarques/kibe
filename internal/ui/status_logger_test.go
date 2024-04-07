package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_logWriterShouldWriteLast10Items(t *testing.T) {
	testCases := [][]byte{
		[]byte("case 0"),
		[]byte("case 1"),
		[]byte("case 2"),
		[]byte("case 3"),
		[]byte("case 4"),
		[]byte("case 5"),
		[]byte("case 6"),
		[]byte("case 7"),
		[]byte("case 8"),
		[]byte("case 9"),
		[]byte("case 10"),
	}

	expected := []string{
		"case 1",
		"case 2",
		"case 3",
		"case 4",
		"case 5",
		"case 6",
		"case 7",
		"case 8",
		"case 9",
		"case 10",
	}

	w := newlogWriter()

	for _, tc := range testCases {
		i, err := w.Write(tc)
		assert.Equal(t, len(tc), i)
		assert.Nil(t, err)
	}

	actual := w.Out().Val().([]string)
	assert.Equal(t, expected, actual)
}
