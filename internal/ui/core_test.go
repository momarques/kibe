package ui

import (
	"testing"

	"github.com/charmbracelet/x/exp/teatest"
)

var (
	coreModel CoreUI = NewUI()
)

func TestFinalModel(t *testing.T) {
	tm := teatest.NewTestModel(t, NewUI(), teatest.WithInitialTermSize(300, 100))
	fm := tm.FinalModel(t)
	m, ok := fm.(CoreUI)
	if !ok {
		t.Fatalf("final model have the wrong type: %T", fm)
	}
	if m.client.Clientset != nil {
		t.Errorf("client not initialized")
	}
}
