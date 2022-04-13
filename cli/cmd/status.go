package cmd

import (
	"bufio"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"io"
	"os"
	"sort"
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

// Display Format: | filename | status | last modification time |
func runStatusCommand(writer io.Writer, trackedFilePath string) error {
	allMetadata, err := getAllDataFromTrackedFile(trackedFilePath)
	if err != nil {
		return err
	}

	displayResults(writer, allMetadata)

	return nil
}

func displayResults(writer io.Writer, allMetadata []fileMetadata) {
	if len(allMetadata) == 0 {
		fmt.Fprintln(writer, "No Changes!")
		return
	}

	// TODO: Status Created = Green, Status Updated = Blue
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"File name", "Status", "Last Modification Time"})
	table.SetRowLine(true)
	// TODO: Modification time format
	for _, mt := range allMetadata {
		table.Append([]string{mt.Name, string(mt.Status), mt.ModificationTime})
		table.SetRowLine(true)
	}
	table.Render()
}

func getAllDataFromTrackedFile(trackedFilePath string) ([]fileMetadata, error) {
	filePtr, err := openFile(trackedFilePath)
	if err != nil {
		return nil, err
	}
	defer filePtr.Close()

	updateLatestStateMap := make(map[string]fileMetadata)

	scanner := bufio.NewScanner(filePtr)
	for scanner.Scan() {
		mData := extractDataFromFile(scanner.Text())
		updateLatestStateMap[mData.Name] = mData
	}

	var allMetadata []fileMetadata
	for _, v := range updateLatestStateMap {
		allMetadata = append(allMetadata, v)
	}

	sort.Slice(allMetadata, func(i, j int) bool {
		return allMetadata[i].Name > allMetadata[j].Name
	})

	return allMetadata, nil
}
