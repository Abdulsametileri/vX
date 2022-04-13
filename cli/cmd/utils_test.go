package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkDirectoryExists(t *testing.T) {
	dirName := ".vx-test"

	t.Run("return false when directory is not exist", func(t *testing.T) {
		exists, err := checkPathExists(dirName)
		assert.Nil(t, err)
		assert.False(t, exists)
	})

	t.Run("return true when directory is exist", func(t *testing.T) {
		err := createDirectory(dirName)
		if err != nil {
			t.Fatal(err)
		}
		defer removeDirectory(dirName)

		exists, err := checkPathExists(dirName)
		assert.True(t, exists)
		assert.Nil(t, err)
	})
}
