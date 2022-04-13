package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var (
	ErrMustSpecifyCommitVersion = errors.New("you must specify commit version")
	ErrCommitVersionNotValid    = errors.New("commit version is not valid")
)

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

var checkoutCmd = &cobra.Command{
	Use:     "checkout",
	Short:   "This allows you to move specified commit version",
	Example: "vx checkout v2",
	RunE: func(_ *cobra.Command, args []string) error {
		return runCheckoutCommand(args)
	},
}

func runCheckoutCommand(args []string) error {
	if err := checkCommitVersionIsSpecified(args); err != nil {
		return err
	}

	commitVersion := args[0]

	if valid, err := isValidCommitVersion(commitVersion); err != nil {
		return err
	} else if !valid {
		return ErrCommitVersionNotValid
	}

	if exist, err := isCheckoutFolderAlreadyExist(commitVersion); err != nil {
		return err
	} else if exist {
		fmt.Println("Checkout file is already exist :)")
		return nil
	}

	checkoutFilePath := filepath.Join(vxCheckoutDirName, commitVersion)
	if err := createDirectory(checkoutFilePath); err != nil {
		return err
	}

	var commands []string

	dirCount, err := getNumberOfChildrenDir(vxCommitDirName)
	if err != nil {
		return err
	}

	for i := 1; i <= dirCount; i++ {
		vNo := fmt.Sprintf("v%d", i)
		cmd := fmt.Sprintf("rsync -a .vx/commit/%s/ %s --exclude=metadata.txt", vNo, checkoutFilePath)
		commands = append(commands, cmd)

		if commitVersion == vNo {
			break
		}
	}

	for _, command := range commands {
		commandStatus, _, stdErr := runShellCommand(command)
		if commandStatus != nil {
			return errors.New(stdErr)
		}
	}

	return nil
}

func checkCommitVersionIsSpecified(args []string) error {
	if len(args) == 0 {
		return ErrMustSpecifyCommitVersion
	}
	return nil
}

func isValidCommitVersion(commitVersion string) (bool, error) {
	dir, err := os.ReadDir(vxCommitDirName)
	if err != nil {
		return false, err
	}

	for _, d := range dir {
		if d.Name() == commitVersion {
			return true, nil
		}
	}

	return false, nil
}

func isCheckoutFolderAlreadyExist(path string) (bool, error) {
	if exist, err := checkPathExists(path); err != nil {
		return false, err
	} else if exist {
		return true, nil
	}
	return false, nil
}
