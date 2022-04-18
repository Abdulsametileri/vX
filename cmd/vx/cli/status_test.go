package cli

import (
	"bytes"
	"testing"

	_ "github.com/Abdulsametileri/vX/testing"
	"github.com/acarl005/stripansi"
	"github.com/stretchr/testify/assert"
)

func Test_runStatusCommand(t *testing.T) {
	var buf bytes.Buffer
	err := runStatusCommand(&buf, "testdata/status.txt")
	assert.Nil(t, err)

	expected := `+-----------------------------+---------+------------------------+
|          FILE NAME          | STATUS  | LAST MODIFICATION TIME |
+-----------------------------+---------+------------------------+
| testdata/z.go               | Created | 2022-04-14 05:11:04    |
| testdata/status.txt         | Created | 2022-04-14 05:42:15    |
| testdata/staging-area.txt   | Created | 2022-04-14 05:42:15    |
| testdata/example/example.go | Created | 2022-04-13 07:41:26    |
| testdata/a2.txt             | Created | 2022-04-13 06:58:03    |
| testdata/a1.txt             | Created | 2022-04-13 06:58:03    |
| README.md                   | Updated | 2022-04-14 05:49:09    |
+-----------------------------+---------+------------------------+
`
	assert.Equal(t, expected, stripansi.Strip(buf.String()))
}
