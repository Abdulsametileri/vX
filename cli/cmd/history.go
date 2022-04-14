package cmd

import (
	"bufio"
	"io"
	"os"
	"path/filepath"

	"github.com/fatih/color"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(historyCmd)
}

var historyCmd = &cobra.Command{
	Use:     "history",
	Short:   "This allows you display your commit history",
	Example: "vx history",
	RunE: func(_ *cobra.Command, _ []string) error {
		return runHistoryCommand(os.Stdout, vxCommitDirPath)
	},
}

func runHistoryCommand(writer io.Writer, commitDirName string) error {
	dirLen, err := getNumberOfChildrenDir(commitDirName)
	if err != nil {
		return err
	}

	history := make([]commitHistory, 0, dirLen)

	// in order to provide desc order, we reverse for variable in the loop.
	for i := dirLen; i > 0; i-- {
		vNo := getDirNameUsingVersion(i)
		metadataFilePtr, _ := openFile(filepath.Join(commitDirName, vNo, vxCommitMetadataFileName))
		scanner := bufio.NewScanner(metadataFilePtr)
		for scanner.Scan() {
			lineStr := scanner.Text()
			ch := extractCommitMetadataFromLine(lineStr)
			ch.Version = vNo
			history = append(history, ch)
		}
	}

	displayHistory(writer, history)

	return nil
}

// Display Format: | Commit Message | Commit Date |
func displayHistory(writer io.Writer, history []commitHistory) {
	if len(history) == 0 {
		color.Green("No commits yet!")
		return
	}

	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"Commit Version (ID)", "Commit Message", "Commit Date"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
	)

	for _, h := range history {
		table.Append([]string{h.Version, h.Message, h.Date})
		table.SetRowLine(true)
	}

	table.Render()
}
