package kube

import (
	"reflect"
	"strings"
	"time"

	"github.com/momarques/kibe/internal/logging"
	"github.com/samber/lo"
)

func DeltaTime(t time.Time) string {
	elapsedTime := time.Since(t)
	elapsedTimeString := elapsedTime.String()

	elapsed, _, _ := strings.Cut(elapsedTimeString, ".")
	return elapsed + "s"
}

func LookupStructFieldNames(t reflect.Type) []string {
	logging.Log.Info("num field ->> ", t.NumField())
	logging.Log.Info("num field ->> ", t.NumField())

	return lo.Times(t.NumField(), func(index int) string {
		logging.Log.Info("field name ->> ", t.Field(index).Name)

		tabName, _ := t.Field(index).Tag.Lookup("kibedescription")
		logging.Log.Info("tag ->> ", tabName)
		return tabName
	})
}
