package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"time"
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

func createCommitMetadataFile(commitDirName, commitMsg string) error {
	msgFilePtr, err := createNestedFile(filepath.Join(commitDirName, "metadata.txt"))
	if err != nil {
		return err
	}
	defer msgFilePtr.Close()

	_, _ = msgFilePtr.WriteString(commitMsg)
	_, _ = msgFilePtr.WriteString("|")
	_, _ = msgFilePtr.WriteString(time.Now().Format(vxTimeFormat))

	return nil
}

func runCommitCommand(trackedFilePath, msg string) error {
	dirCount, err := getNumberOfChildrenDir(vxCommitDirName)
	if err != nil {
		return err
	}

	newCommitDirName := filepath.Join(vxCommitDirName, fmt.Sprintf("v%d", dirCount+1))
	err = createCommitMetadataFile(newCommitDirName, msg)
	if err != nil {
		return err
	}

	fileNameToMetadata, err := createFileNameToMetadataMap(trackedFilePath)
	if err != nil {
		return err
	}

	for _, file := range fileNameToMetadata {
		firstCommitFilePath := filepath.Join(vxFirstCommitDirName, file.Name)
		destCommitFilePath := filepath.Join(newCommitDirName, file.Name)

		exist, _ := checkPathExists(firstCommitFilePath)
		if exist {
			// TODO: Instead of new copy of the fresh file, we can copy only changes and apply them
		}
		destinationFilePtr, _ := createNestedFile(destCommitFilePath)
		originalFilePtr, _ := os.Open(file.Name)
		_, _ = io.Copy(destinationFilePtr, originalFilePtr)

		originalFilePtr.Close()
		destinationFilePtr.Close()

	}

	stagingFilePtr, _ := openFile(stagingAreaFilePath)
	defer stagingFilePtr.Close()

	err = clearFileContent(stagingFilePtr)

	return err
}
