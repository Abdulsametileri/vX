package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkCommitVersionIsSpecified(t *testing.T) {
	err := checkCommitVersionIsSpecified([]string{})
	assert.ErrorIs(t, err, ErrMustSpecifyCommitVersion)
}
