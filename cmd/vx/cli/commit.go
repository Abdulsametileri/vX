package cli

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

const (
	messageFlag = "message"
)

func init() {
	commitCmd.Flags().StringP(messageFlag, "m", "", "the commit message")
	_ = commitCmd.MarkFlagRequired(messageFlag)
	rootCmd.AddCommand(commitCmd)
}

var commitCmd = &cobra.Command{
	Use:     "commit",
	Short:   "This allows you save all all file changes",
	Example: `vx commit -m "your message"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		msg, _ := cmd.Flags().GetString(messageFlag)
		if msg == "" {
			return errors.New("commit message cannot be empty")
		}
		return runCommitCommand(stagingAreaFilePath, msg)
	},
}

func runCommitCommand(trackedFilePath, msg string) error {
	dirCount, err := getNumberOfChildrenDir(vxCommitDirPath)
	if err != nil {
		return err
	}

	newCommitDirName := filepath.Join(vxCommitDirPath, fmt.Sprintf("v%d", dirCount+1))
	err = createCommitMetadataFile(newCommitDirName, msg, time.Now())
	if err != nil {
		return err
	}

	fileNameToMetadata, err := createFileNameToMetadataMap(trackedFilePath)
	if err != nil {
		return err
	}

	for _, file := range fileNameToMetadata {
		destCommitFilePath := filepath.Join(newCommitDirName, file.Path)

		destinationFilePtr, _ := createNestedFile(destCommitFilePath)
		originalFilePtr, _ := os.Open(file.Path)
		_, _ = io.Copy(destinationFilePtr, originalFilePtr)

		originalFilePtr.Close()
		destinationFilePtr.Close()
	}

	stagingFilePtr, _ := openFile(stagingAreaFilePath)
	defer stagingFilePtr.Close()

	err = clearFileContent(stagingFilePtr)

	return err
}

func createCommitMetadataFile(commitDirName, commitMsg string, commitDate time.Time) error {
	msgFilePtr, err := createNestedFile(filepath.Join(commitDirName, vxCommitMetadataFileName))
	if err != nil {
		return err
	}
	defer msgFilePtr.Close()

	_, _ = msgFilePtr.WriteString(toFormatCommitMetadata(commitMsg, commitDate))

	return nil
}
