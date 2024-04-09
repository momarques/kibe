package kube

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/samber/lo"
)

func DeltaTime(t1, t2 time.Time) string {
	elapsedTime := t2.Sub(t1)
	if elapsedTime.Hours() > 24 {
		return fmt.Sprintf("%dd%dh", int(elapsedTime.Hours()/24), int(elapsedTime.Hours())%24)
	}
	elapsedTimeString := elapsedTime.String()
	elapsed, _, _ := strings.Cut(elapsedTimeString, ".")
	return elapsed
}

func LookupStructFieldNames(s any) []string {
	t := reflect.TypeOf(s)
	return lo.Times(t.NumField(),
		func(index int) string {
			tabName, _ := t.Field(index).Tag.Lookup("kibedescription")
			return tabName
		})
}
