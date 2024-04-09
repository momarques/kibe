package kube

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_DeltaTime(t *testing.T) {
	testCases := []time.Time{
		time.Date(2024, time.February, 29, 23, 57, 30, 0, time.Local),
		time.Date(2024, time.February, 29, 16, 35, 30, 0, time.Local),
		time.Date(2024, time.February, 12, 16, 35, 30, 0, time.Local),
		time.Date(2024, time.February, 12, 16, 33, 30, 0, time.Local),
		time.Date(2024, time.February, 12, 15, 34, 30, 0, time.Local),
		time.Date(2024, time.February, 11, 16, 34, 30, 0, time.Local),
		time.Date(2024, time.February, 5, 16, 34, 30, 0, time.Local),
		time.Date(2024, time.January, 5, 16, 34, 30, 0, time.Local),
		time.Date(2023, time.January, 5, 16, 34, 30, 0, time.Local),
	}
	referenceDate := time.Date(2024, time.March, 1, 0, 0, 0, 0, time.Local)

	expectedStrings := []string{
		"2m30s",
		"7h24m30s",
		"17d7h",
		"17d7h",
		"17d8h",
		"18d7h",
		"24d7h",
		"55d7h",
		"420d7h",
	}

	for i, duration := range testCases {
		actual := DeltaTime(duration, referenceDate)
		if !assert.Equal(t, expectedStrings[i], actual) {
			t.Logf("test case n%d -> expected: %s, got: %s\n", i, expectedStrings[i], actual)
		}
	}
}

func Test_LookupStructFieldNames(t *testing.T) {
	type S struct {
		FullName   string `kibedescription:"Full Name"`
		MotherName string `kibedescription:"Mother's Name"`
		Age        int    `kibedescription:"Age"`
	}

	s := S{
		FullName:   "John Doe",
		MotherName: "Jane Doe",
		Age:        30,
	}

	expected := []string{"Full Name", "Mother's Name", "Age"}
	actual := LookupStructFieldNames(s)
	assert.Equal(t, expected, actual)
}
