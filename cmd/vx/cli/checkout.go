package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
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
		return runCheckoutCommand(vxCheckoutDirPath, vxCommitDirPath, args)
	},
}

func runCheckoutCommand(checkoutDirPath, commitDirPath string, args []string) error {
	if err := checkCommitVersionIsSpecified(args); err != nil {
		return err
	}

	commitVersion := args[0]

	if valid, err := isValidCommitVersion(commitDirPath, commitVersion); err != nil {
		return err
	} else if !valid {
		return ErrCommitVersionNotValid
	}

	checkoutFilePath := filepath.Join(checkoutDirPath, commitVersion)

	if exist, err := isCheckoutFolderAlreadyExist(checkoutFilePath); err != nil {
		return err
	} else if exist { // checkout version folder already exist no need to do operation
		return nil
	}

	if err := createDirectory(checkoutFilePath); err != nil {
		return err
	}

	var commands []string

	dirCount, err := getNumberOfChildrenDir(commitDirPath)
	if err != nil {
		return err
	}

	// Copy and merge all the changes from first version to given commit version
	for i := 1; i <= dirCount; i++ {
		vNo := getDirNameUsingVersion(i)
		cmd := fmt.Sprintf("rsync -a .vx/commit/%s/ %s --exclude=metadata.txt", vNo, checkoutFilePath)
		commands = append(commands, cmd)

		if commitVersion == vNo {
			break
		}
	}

	for _, command := range commands {
		stdErr, _, commandStatus := runShellCommand(command)
		if commandStatus != nil {
			return errors.New(stdErr)
		}
	}

	color.Green("All files merged in %s directory!", checkoutFilePath)

	return nil
}

func checkCommitVersionIsSpecified(args []string) error {
	if len(args) == 0 {
		return ErrMustSpecifyCommitVersion
	}
	return nil
}

func isValidCommitVersion(commitDirPath, commitVersion string) (bool, error) {
	dir, err := os.ReadDir(commitDirPath)
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
