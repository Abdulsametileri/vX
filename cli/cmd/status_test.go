package cmd

import (
	"os"
	"testing"
)

func Test_runStatusCommand(t *testing.T) {
	runStatusCommand(os.Stdout, ".vx/trackedFiles.txt")
}
