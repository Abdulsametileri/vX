package cmd

import (
	"fmt"
	"testing"
)

func Test_runCommitCommand(t *testing.T) {
	err := runCommitCommand(stagingAreaFilePath, "first commit")
	fmt.Println(err)
}
