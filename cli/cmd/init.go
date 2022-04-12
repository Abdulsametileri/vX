package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

const (
	vxRootDirName = ".vx"
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
	exists, err := checkDirectoryExists(vxRootDirName)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("vx root directory is already exist")
	}

	err = createDirectory(vxRootDirName)
	if err != nil {
		return err
	}
	return nil
}
