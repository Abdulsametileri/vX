package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testTrackedFile = "testdata/test_tracked.txt"
)

func Test_runAddCommand(t *testing.T) {
	files := []string{"testdata/z.go"}
	err := runAddCommand(testTrackedFile, files)
	assert.Nil(t, err)
}

func Test_CreateFileNameToModificationMap(t *testing.T) {
	expectedMap := fileNameToMetadataMap{
		"go.mod": fileMetadata{
			Status:           "Created",
			ModificationTime: "2022-04-12 12:35:01.344508354 +0300 +03",
		},
		"README.md": fileMetadata{
			Status:           "Created",
			ModificationTime: "2022-04-12 12:51:16.579300203 +0300 +03",
		},
	}

	m, err := createFileNameToMetadataMap(testTrackedFile)
	assert.Nil(t, err)
	assert.Equal(t, expectedMap, m)
}
