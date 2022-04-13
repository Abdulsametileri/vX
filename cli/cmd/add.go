package cmd

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type fileNameToMetadataMap map[string]fileMetadata

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "This allows you to track all file status (created, modified, deleted)",
	Example: "vx add main.go images/",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAddCommand(stagingAreaFilePath, args)
	},
}

func runAddCommand(stagingAreaFilePath string, filePaths []string) error {
	fileNameToMetadata, err := createFileNameToMetadataMap(statusFilePath)
	if err != nil {
		return err
	}

	for _, filePath := range filePaths {
		fileInfo, _ := os.Stat(filePath)

		if fileInfo.IsDir() {
			traverseDirectory(filePath, fileNameToMetadata, fileInfo)
		} else {
			extractFileMetadata(fileNameToMetadata, filePath, fileInfo.ModTime().String())
		}
	}

	statusFilePtr, _ := openFile(statusFilePath)
	defer statusFilePtr.Close()

	stagingFilePtr, _ := openFileAppendMode(stagingAreaFilePath)
	defer stagingFilePtr.Close()

	err = clearFileContent(statusFilePtr)
	if err != nil {
		return err
	}

	for fileName, metadata := range fileNameToMetadata {
		lineStr := fmt.Sprintf("%s|%s|%s\n", fileName,
			metadata.ModificationTime, string(metadata.Status))
		statusFilePtr.WriteString(lineStr)
		if metadata.GoToStaging {
			stagingFilePtr.WriteString(lineStr)
		}
	}

	return nil
}

func traverseDirectory(arg string, fileNameToMetadata fileNameToMetadataMap, fileInfo os.FileInfo) {
	_ = filepath.Walk(arg, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			extractFileMetadata(fileNameToMetadata, path, fileInfo.ModTime().String())
		}
		return nil
	})
}

func extractFileMetadata(fileNameToMetadata fileNameToMetadataMap, filePath, fileCurrModTime string) {
	metadata, ok := fileNameToMetadata[filePath]
	if ok {
		if fileCurrModTime != metadata.ModificationTime {
			fileNameToMetadata[filePath] = fileMetadata{
				Name:             filePath,
				Status:           StatusUpdated,
				ModificationTime: fileCurrModTime,
				GoToStaging:      true,
			}
		}
	} else {
		fileNameToMetadata[filePath] = fileMetadata{
			Name:             filePath,
			Status:           StatusCreated,
			ModificationTime: fileCurrModTime,
			GoToStaging:      true,
		}
	}
}

func createFileNameToMetadataMap(filePath string) (fileNameToMetadataMap, error) {
	fileNameToModification := make(fileNameToMetadataMap)

	trackedFilePtr, err := openFile(filePath)
	if err != nil {
		return fileNameToModification, err
	}
	defer trackedFilePtr.Close()

	scanner := bufio.NewScanner(trackedFilePtr)

	for scanner.Scan() {
		line := scanner.Text()
		fm := extractDataFromFile(line)
		fileNameToModification[fm.Name] = fm
	}

	err = scanner.Err()
	if err != nil {
		return fileNameToModification, err
	}

	return fileNameToModification, nil
}
