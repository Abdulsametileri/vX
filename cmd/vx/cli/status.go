package cli

import (
	"bufio"
	"io"
	"os"
	"sort"

	"github.com/fatih/color"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "This allows you to display all tracked files status",
	Example: "vx status",
	RunE: func(_ *cobra.Command, _ []string) error {
		return runStatusCommand(os.Stdout, stagingAreaFilePath)
	},
}

func runStatusCommand(writer io.Writer, stagingAreaFilePath string) error {
	allMetadata, err := getFileMetadataFromStagingFile(stagingAreaFilePath)
	if err != nil {
		return err
	}

	displayStatus(writer, allMetadata)

	return nil
}

// Display Format: | filename | status | last modification time |
func displayStatus(writer io.Writer, allMetadata []fileMetadata) {
	if len(allMetadata) == 0 {
		color.Green("No changes on staging area!")
		return
	}

	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"File name", "Status", "Last Modification Time"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
	)

	for _, mt := range allMetadata {
		row := []string{mt.Path, string(mt.Status), mt.ModificationTime}
		statusColor := tablewriter.FgGreenColor
		if mt.Status == StatusUpdated {
			statusColor = tablewriter.FgBlueColor
		}
		table.Rich(row, []tablewriter.Colors{{}, {tablewriter.Bold, statusColor}, {}})
	}

	table.Render()
}

func getFileMetadataFromStagingFile(stagingAreaFilePath string) (fileMetadataArr, error) {
	filePtr, err := openFile(stagingAreaFilePath)
	if err != nil {
		return nil, err
	}
	defer filePtr.Close()

	// Because of the staging is a type of append-only file,
	// we have to update file status to latest
	updateLatestStateMap := make(map[string]fileMetadata)

	scanner := bufio.NewScanner(filePtr)
	for scanner.Scan() {
		mData := extractFileMetadataFromLine(scanner.Text())
		updateLatestStateMap[mData.Path] = mData
	}

	var allMetadata fileMetadataArr
	for _, v := range updateLatestStateMap {
		allMetadata = append(allMetadata, v)
	}

	sort.Sort(allMetadata)

	return allMetadata, nil
}
