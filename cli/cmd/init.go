package cmd

import (
	"errors"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	statusFilePath      = filepath.Join(vxRootDirName, "status.txt")
	stagingAreaFilePath = filepath.Join(vxRootDirName, "staging-area.txt")
)

const (
	vxRootDirName            = ".vx"
	vxCommitDirPath          = ".vx/commit"
	vxCheckoutDirPath        = ".vx/checkout"
	vxStatusFilePath         = ".vx/status.txt"
	vxStagingFilePath        = ".vx/staging-area.txt"
	vxTimeFormat             = "2006-01-02 03:04:05"
	vxCommitMetadataFileName = "metadata.txt"
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
	if exists, err := checkPathExists(vxRootDirName); err != nil {
		return err
	} else if exists {
		return errors.New("vx root directory is already exist")
	}

	if err := createDirectories(vxRootDirName, vxCommitDirPath, vxCheckoutDirPath); err != nil {
		return err
	}

	if err := createFile(vxStatusFilePath); err != nil {
		return err
	}

	if err := createFile(vxStagingFilePath); err != nil {
		return err
	}

	color.Green("All files are initialized within .vx/ directory!")

	return nil
}
