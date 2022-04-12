package cmd

import (
	"bufio"
	"io"
	"os"
	"sort"

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
		return runStatusCommand(os.Stdout, trackedFilesPath)
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
	// TODO: Status Created = Green, Status Updated = Blue
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"File name", "Status", "Last Modification Time"})
	table.SetRowLine(true)
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

	var allMetadata []fileMetadata

	scanner := bufio.NewScanner(filePtr)
	for scanner.Scan() {
		allMetadata = append(allMetadata, extractDataFromFile(scanner.Text()))
	}

	sort.Slice(allMetadata, func(i, j int) bool {
		return allMetadata[i].Name > allMetadata[j].Name
	})

	return allMetadata, nil
}
