package pod

import (
	"testing"
	"time"
)

func Test_DeltaTime(t *testing.T) {
	var testCases = []time.Time{
		time.Date(2024, time.February, 12, 16, 35, 30, 0, time.Local),
		time.Date(2024, time.February, 12, 16, 33, 30, 0, time.Local),
		time.Date(2024, time.February, 12, 15, 34, 30, 0, time.Local),
		time.Date(2024, time.February, 11, 16, 34, 30, 0, time.Local),
		time.Date(2024, time.February, 5, 16, 34, 30, 0, time.Local),
		time.Date(2024, time.January, 5, 16, 34, 30, 0, time.Local),
		time.Date(2023, time.January, 5, 16, 34, 30, 0, time.Local),
	}

	for i, duration := range testCases {
		t.Logf("test case n%d -> %s\n", i, DeltaTime(duration))
	}
	t.Fail()
}
