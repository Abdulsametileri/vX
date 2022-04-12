package cmd

import (
	"io"
	"os"
)

func checkDirectoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func createDirectory(dirName string) error {
	return os.Mkdir(dirName, os.ModePerm)
}

func openFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_CREATE|os.O_RDWR, os.ModePerm)
}

func clearFileContent(filePtr *os.File) error {
	err := filePtr.Truncate(0)
	filePtr.Seek(0, io.SeekStart) //nolint:errcheck
	return err
}

func removeDirectory(dirName string) error {
	return os.RemoveAll(dirName)
}
