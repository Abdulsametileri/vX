package cmd

import (
	"fmt"
	"strings"
)

const (
	separator = "|"
)

func (f *fileMetadata) toFormatFileMetadataForFile() string {
	return fmt.Sprintf(
		"%s%s%s%s%s\n",
		f.Path,
		separator,
		f.ModificationTime,
		separator,
		string(f.Status),
	)
}

func getDirNameUsingVersion(versionNo int) string {
	return fmt.Sprintf("v%d", versionNo)
}

func extractFileMetadataFromLine(lineStr string) fileMetadata {
	structure := strings.Split(lineStr, separator)
	return fileMetadata{
		Path:             structure[0],
		ModificationTime: structure[1],
		Status:           FileStatus(structure[2]),
	}
}

func extractCommitMetadataFromLine(lineStr string) commitHistory {
	structure := strings.Split(lineStr, separator)
	return commitHistory{
		Message: structure[0],
		Date:    structure[1],
	}
}
