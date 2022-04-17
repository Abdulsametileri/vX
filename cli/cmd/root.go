package cmd

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	changeWorkingDirectoryToRootDir()
	changeGlobalTimeToUTC()
}

var UseDefaultTime bool

var rootCmd = &cobra.Command{
	Use:   "vx",
	Short: "vX is a Command Line Interface (CLI) to implement basic version control system.",
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&UseDefaultTime, "dtime", false, "use default time (0001-01-01 00:00:00 +0000 UTC)")
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

func changeGlobalTimeToUTC() {
	loc, err := time.LoadLocation("UTC")
	if err == nil {
		time.Local = loc
	}
}
