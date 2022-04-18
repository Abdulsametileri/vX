package cli

import (
	"testing"

	_ "github.com/Abdulsametileri/vX/testing"
	"github.com/stretchr/testify/assert"
)

func Test_CreateFileNameToModificationMap(t *testing.T) {
	lines := []string{
		"testdata/a2.txt|2022-04-13 06:58:03|Created",
		"testdata/example/example.go|2022-04-13 07:41:26|Created",
		"testdata/staging-area.txt|2022-04-14 05:42:15|Created",
		"testdata/status.txt|2022-04-14 05:42:15|Created",
		"testdata/z.go|2022-04-14 05:11:04|Created",
		"README.md|2022-04-14 05:42:11|Created",
		"testdata/a1.txt|2022-04-13 06:58:03|Created",
		"README.md|2022-04-14 05:49:09|Updated",
	}
	expectedMap := make(filePathToMetadataMap)
	for _, line := range lines {
		fm := extractFileMetadataFromLine(line)
		expectedMap[fm.Path] = fm
	}

	m, err := createFileNameToMetadataMap("testdata/staging-area.txt")
	assert.Nil(t, err)
	assert.Equal(t, expectedMap, m)
}
