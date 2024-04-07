package ui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
)

func Test_applyStatusBarChanges(t *testing.T) {
	var cmd tea.Cmd

	expected := newStatusBarModel()
	expected.SecondColumn = "Pod"
	expected.ThirdColumn = "Context: context-kibe"
	expected.FourthColumn = "Namespace: default-ns"

	actual := coreModel
	actual, cmd = actual.applyStatusBarChanges(statusBarUpdated{"Pod", "context-kibe", "default-ns"})

	assert.Nil(t, cmd)
	assert.Equal(t, expected, actual.statusBar)
}

func Test_updateStatusBarMsgGeneration(t *testing.T) {
	expected := statusBarUpdated{"Pod", "context-kibe", "default-ns"}
	actual := updateStatusBar("Pod", "context-kibe", "default-ns")

	assert.Equal(t, expected, actual())
}
