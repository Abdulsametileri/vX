package cli

import (
	"bytes"
	"fmt"
	"testing"

	_ "github.com/Abdulsametileri/vX/testing"
	"github.com/acarl005/stripansi"
	"github.com/stretchr/testify/assert"
)

func Test_runHistoryCommand(t *testing.T) {
	var buf bytes.Buffer
	err := runHistoryCommand(&buf, "testdata/commit")
	expected := `+---------------------+----------------+---------------------+
| COMMIT VERSION (ID) | COMMIT MESSAGE |     COMMIT DATE     |
+---------------------+----------------+---------------------+
| v2                  | second commit  | 2022-04-14 07:30:21 |
+---------------------+----------------+---------------------+
| v1                  | first commit   | 2022-04-14 07:30:05 |
+---------------------+----------------+---------------------+
`

	fmt.Println(stripansi.Strip(buf.String()))
	assert.Nil(t, err)
	assert.Equal(t, expected, stripansi.Strip(buf.String()))
}
