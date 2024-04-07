package kube

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PodDescriptionTabNames(t *testing.T) {
	expected := []string{
		"Overview",
		"Status",
		"Labels",
		"Annotations",
		"Volumes",
		"Containers",
		"Scheduling",
		"Node Scheduling",
		"Events",
	}

	podDesc := new(PodDescription)
	for _, p := range podDesc.TabNames() {
		assert.Contains(t, expected, p)
	}
}
