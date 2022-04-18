package testing

import (
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	changeWorkingDirectoryToRootDir()
}

func changeWorkingDirectoryToRootDir() {
	_, filename, _, _ := runtime.Caller(0)
	err := os.Chdir(filepath.Join(filepath.Dir(filename), ".."))
	if err != nil {
		panic(err)
	}
}
