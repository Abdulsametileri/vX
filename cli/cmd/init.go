package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"path/filepath"
)

var (
	statusFilePath      = filepath.Join(vxRootDirName, "status.txt")
	stagingAreaFilePath = filepath.Join(vxRootDirName, "staging-area.txt")
)

const (
	vxRootDirName        = ".vx"
	vxCommitDirName      = ".vx/commit"
	vxCheckoutDirName    = ".vx/checkout"
	vxFirstCommitDirName = ".vx/commit/v1"
	vxStatusFileName     = ".vx/status.txt"
	vxStagingFileName    = ".vx/staging-area.txt"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "This allows you initialize vX control system",
	Example: "vx init",
	RunE: func(_ *cobra.Command, _ []string) error {
		return runInitCommand()
	},
}

func runInitCommand() error {
	exists, err := checkPathExists(vxRootDirName)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("vx root directory is already exist")
	}

	if err = createDirectory(vxRootDirName); err != nil {
		return err
	}

	if err = createDirectory(vxCommitDirName); err != nil {
		return err
	}

	if err = createFile(vxStatusFileName); err != nil {
		return err
	}

	if err = createFile(vxStagingFileName); err != nil {
		return err
	}

	return nil
}
