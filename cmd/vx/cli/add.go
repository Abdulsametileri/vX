package cli

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

type filePathToMetadataMap map[string]fileMetadata

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "This allows you to track all file status (created, modified)",
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

	determineStagingAreaFiles(filePaths, fileNameToMetadata)

	statusFilePtr, _ := openFile(statusFilePath)
	defer statusFilePtr.Close()

	stagingFilePtr, _ := openFileAppendMode(stagingAreaFilePath)
	defer stagingFilePtr.Close()

	// we need to truncate status folder in order to catch overwrite file status
	if err := clearFileContent(statusFilePtr); err != nil {
		return err
	}

	for _, metadata := range fileNameToMetadata {
		lineStr := metadata.toFormatFileMetadataForFile()

		_, _ = statusFilePtr.WriteString(lineStr)
		if metadata.GoToStaging {
			_, _ = stagingFilePtr.WriteString(lineStr)
		}
	}

	return nil
}

func determineStagingAreaFiles(filePaths []string, fileNameToMetadata filePathToMetadataMap) {
	for _, filePath := range filePaths {
		fileInfo, _ := os.Stat(filePath)

		if fileInfo.IsDir() {
			traverseDirectory(filePath, fileNameToMetadata)
		} else {
			extractFileMetadata(fileNameToMetadata, filePath, fileInfo.ModTime())
		}
	}
}

func traverseDirectory(arg string, fileNameToMetadata filePathToMetadataMap) {
	_ = filepath.Walk(arg, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			extractFileMetadata(fileNameToMetadata, path, info.ModTime())
		}
		return nil
	})
}

func extractFileMetadata(fileNameToMetadata filePathToMetadataMap, filePath string, fileCurrModTime time.Time) {
	fileCurModTimeStr := fileCurrModTime.Format(vxTimeFormat)
	metadata, ok := fileNameToMetadata[filePath]

	mStruct := fileMetadata{
		Path:             filePath,
		ModificationTime: fileCurModTimeStr,
		GoToStaging:      true,
		Status:           StatusCreated,
	}

	if ok && fileCurModTimeStr != metadata.ModificationTime {
		mStruct.Status = StatusUpdated
	}

	fileNameToMetadata[filePath] = mStruct
}

func createFileNameToMetadataMap(filePath string) (filePathToMetadataMap, error) {
	trackedFilePtr, err := openFile(filePath)
	if err != nil {
		return nil, err
	}
	defer trackedFilePtr.Close()

	fileNameToModification := make(filePathToMetadataMap)

	scanner := bufio.NewScanner(trackedFilePtr)
	for scanner.Scan() {
		line := scanner.Text()
		fm := extractFileMetadataFromLine(line)
		fileNameToModification[fm.Path] = fm
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return fileNameToModification, nil
}
