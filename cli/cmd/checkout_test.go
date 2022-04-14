package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkCommitVersionIsSpecified(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		err := checkCommitVersionIsSpecified([]string{})
		assert.ErrorIs(t, err, ErrMustSpecifyCommitVersion)
	})
	t.Run("success", func(t *testing.T) {
		err := checkCommitVersionIsSpecified([]string{"v1"})
		assert.Nil(t, err)
	})
}

func Test_isValidCommitVersion(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		valid, err := isValidCommitVersion("testdata/commit", "v1")
		assert.Nil(t, err)
		assert.True(t, valid)
	})
	t.Run("invalid", func(t *testing.T) {
		valid, err := isValidCommitVersion("testdata/commit", "v4")
		assert.Nil(t, err)
		assert.False(t, valid)
	})
}

func Test_isCheckoutFolderAlreadyExist(t *testing.T) {
	t.Run("exist", func(t *testing.T) {
		exist, err := isCheckoutFolderAlreadyExist("testdata/checkout/v1")
		assert.Nil(t, err)
		assert.True(t, exist)
	})
	t.Run("does not exist", func(t *testing.T) {
		exist, err := isCheckoutFolderAlreadyExist("testdata/checkout/v10")
		assert.Nil(t, err)
		assert.False(t, exist)
	})
}

func Test_runCheckoutCommand(t *testing.T) {
	runCheckoutCommand(vxCheckoutDirPath, vxCommitDirPath, []string{"v1"})
}
