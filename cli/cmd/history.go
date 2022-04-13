package cmd

import (
	"bufio"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type commitHistory struct {
	Date    string
	Message string
}

func init() {
	rootCmd.AddCommand(historyCmd)
}

var historyCmd = &cobra.Command{
	Use:     "history",
	Short:   "This allows you display your commit history",
	Example: "vx history",
	RunE: func(_ *cobra.Command, _ []string) error {
		return runHistoryCommand()
	},
}

func runHistoryCommand() error {
	dirLen, err := getNumberOfChildrenDir(vxCommitDirName)
	if err != nil {
		return err
	}

	history := make([]commitHistory, 0, dirLen)

	for i := dirLen; i > 0; i-- {
		vNo := fmt.Sprintf("v%d", i)
		metadataFilePtr, _ := openFile(filepath.Join(vxCommitDirName, vNo, "metadata.txt"))
		scanner := bufio.NewScanner(metadataFilePtr)
		for scanner.Scan() {
			structure := strings.Split(scanner.Text(), "|")
			commitMsg := structure[0]
			commitDate := structure[1]

			history = append(history, commitHistory{
				Date:    commitDate,
				Message: commitMsg,
			})
			fmt.Println(commitDate)
		}
	}

	displayHistory(os.Stdout, history)

	return nil
}

func displayHistory(writer io.Writer, history []commitHistory) {
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"Commit Message", "Commit Date"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlueColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlueColor},
	)

	for _, h := range history {
		table.Append([]string{h.Message, h.Date})
		table.SetRowLine(true)
	}
	table.Render()
}
