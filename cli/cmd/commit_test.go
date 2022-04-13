package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_runCommitCommand(t *testing.T) {
	err := runCommitCommand(stagingAreaFilePath, "first commit")
	assert.Nil(t, err)
}
