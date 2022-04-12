package cmd

import (
	"bufio"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	trackedFilesPath = filepath.Join(vxRootDirName, "status.txt")
	separator        = "|"
	newLine          = "\n"
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
		return runAddCommand(trackedFilesPath, args)
	},
}

func runAddCommand(filePath string, args []string) error {
	trackedFilePtr, err := openFile(filePath)
	if err != nil {
		return err
	}
	defer trackedFilePtr.Close()

	fileNameToMetadata, err := createFileNameToMetadataMap(trackedFilePtr)
	if err != nil {
		return err
	}

	for _, arg := range args {
		fileInfo, _ := os.Stat(arg)

		if fileInfo.IsDir() {
			traverseDirectory(arg, fileNameToMetadata, fileInfo)
		} else {
			extractFileMetadata(fileNameToMetadata, fileInfo.Name(), fileInfo.ModTime().String())
		}
	}

	err = clearFileContent(trackedFilePtr)
	if err != nil {
		return err
	}

	writeContentsToFile(trackedFilePtr, fileNameToMetadata)

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

func extractFileMetadata(fileNameToMetadata fileNameToMetadataMap, fileName, fileCurrModTime string) {
	metadata, ok := fileNameToMetadata[fileName]
	if ok {
		if fileCurrModTime != metadata.ModificationTime {
			fileNameToMetadata[fileName] = fileMetadata{
				Name:             fileName,
				Status:           StatusUpdated,
				ModificationTime: fileCurrModTime,
			}
		}
	} else {
		fileNameToMetadata[fileName] = fileMetadata{
			Name:             fileName,
			Status:           StatusCreated,
			ModificationTime: fileCurrModTime,
		}
	}
}

func createFileNameToMetadataMap(reader io.Reader) (fileNameToMetadataMap, error) {
	scanner := bufio.NewScanner(reader)

	fileNameToModification := make(fileNameToMetadataMap)

	for scanner.Scan() {
		line := scanner.Text()
		fm := extractDataFromFile(line)
		fileNameToModification[fm.Name] = fm
	}

	err := scanner.Err()
	if err != nil {
		return fileNameToModification, err
	}

	return fileNameToModification, nil
}

//nolint:errcheck
func writeContentsToFile(filePtr *os.File, fileNameToMetadata fileNameToMetadataMap) {
	for fileName, metadata := range fileNameToMetadata {
		filePtr.WriteString(fileName)
		filePtr.WriteString(separator)
		filePtr.WriteString(metadata.ModificationTime)
		filePtr.WriteString(separator)
		filePtr.WriteString(string(metadata.Status))
		filePtr.WriteString(newLine)
	}
}
