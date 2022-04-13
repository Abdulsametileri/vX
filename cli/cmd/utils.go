package cmd

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func checkPathExists(path string) (bool, error) {
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

func createFile(name string) error {
	_, err := openFile(name)
	return err
}

func openFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_CREATE|os.O_RDWR, os.ModePerm)
}

func openFileAppendMode(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
}

func clearFileContent(filePtr *os.File) error {
	err := filePtr.Truncate(0)
	filePtr.Seek(0, io.SeekStart) //nolint:errcheck
	return err
}

func removeDirectory(dirName string) error {
	return os.RemoveAll(dirName)
}

func createNestedFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), os.ModePerm); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func getNumberOfChildrenDir(path string) (int, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return 0, err
	}
	return len(files), nil
}
