package kube

import (
	"reflect"
	"strings"
	"time"

	"github.com/samber/lo"
)

func DeltaTime(t time.Time) string {
	elapsedTime := time.Since(t)
	elapsedTimeString := elapsedTime.String()

	elapsed, _, _ := strings.Cut(elapsedTimeString, ".")
	return elapsed + "s"
}

func LookupStructFieldNames(t reflect.Type) []string {
	return lo.Times(t.NumField(),
		func(index int) string {
			tabName, _ := t.Field(index).Tag.Lookup("kibedescription")
			return tabName
		})
}
