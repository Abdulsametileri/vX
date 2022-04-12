package cmd

import "strings"

type FileStatus string

var (
	StatusCreated FileStatus = "Created"
	StatusUpdated FileStatus = "Updated"
)

type fileMetadata struct {
	Name             string
	Status           FileStatus
	ModificationTime string
}

func extractDataFromFile(lineStr string) fileMetadata {
	structure := strings.Split(lineStr, separator)
	return fileMetadata{
		Name:             structure[0],
		ModificationTime: structure[1],
		Status:           FileStatus(structure[2]),
	}
}
