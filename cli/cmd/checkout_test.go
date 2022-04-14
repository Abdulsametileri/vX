package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkCommitVersionIsSpecified(t *testing.T) {
	err := checkCommitVersionIsSpecified([]string{})
	assert.ErrorIs(t, err, ErrMustSpecifyCommitVersion)
}
