package cmd

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

func init() {
	changeWorkingDirectoryToRootDir()
}

var rootCmd = &cobra.Command{
	Use:   "vx",
	Short: "vX is a Command Line Interface (CLI) to implement basic version control system.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute command. Reason: %#v", err)
	}
}

func changeWorkingDirectoryToRootDir() {
	_, filename, _, _ := runtime.Caller(0)
	err := os.Chdir(filepath.Join(filepath.Dir(filename), "..", ".."))
	if err != nil {
		panic(err)
	}
}
