package cmd

import (
	"bytes"
	"testing"

	"github.com/acarl005/stripansi"
	"github.com/stretchr/testify/assert"
)

func Test_runHistoryCommand(t *testing.T) {
	var buf bytes.Buffer
	err := runHistoryCommand(&buf, "testdata/commit")
	expected := `+---------------------+----------------+---------------------+
| COMMIT VERSION (ID) | COMMIT MESSAGE |     COMMIT DATE     |
+---------------------+----------------+---------------------+
| v3                  | third commit   | 2022-04-14 07:30:40 |
+---------------------+----------------+---------------------+
| v2                  | second commit  | 2022-04-14 07:30:21 |
+---------------------+----------------+---------------------+
| v1                  | first commit   | 2022-04-14 07:30:05 |
+---------------------+----------------+---------------------+
`

	assert.Nil(t, err)
	assert.Equal(t, expected, stripansi.Strip(buf.String()))
}
