package ui

import (
	"testing"
	"time"

	"github.com/samber/lo"
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

func Test_formatLogAttrValue(t *testing.T) {
	duration := time.Date(2024, time.February, 29, 0, 0, 0, 0, time.Local).Sub(
		time.Date(2024, time.February, 28, 0, 0, 0, 0, time.Local),
	)
	testCases := []interface{}{
		"teste1",
		duration,
		1,
	}
	expected := []string{
		"teste1",
		"24h0m0sms",
		"1",
	}

	actual := lo.Map(testCases,
		func(item interface{}, _ int) string {
			return formatLogAttrValue(item)
		})

	assert.Equal(t, expected, actual)
}
