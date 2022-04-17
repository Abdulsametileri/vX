package cmd

import (
	"fmt"
	"strings"
	"time"
)

const (
	separator   = "|"
	defaultTime = "0001-01-01 00:00:00 +0000 UTC"
)

func (f *fileMetadata) toFormatFileMetadataForFile() string {
	mTime := f.ModificationTime
	if UseDefaultTime {
		mTime = defaultTime
	}

	return fmt.Sprintf(
		"%s%s%s%s%s\n",
		f.Path,
		separator,
		mTime,
		separator,
		string(f.Status),
	)
}

func toFormatCommitMetadata(commitMsg string, commitDate time.Time) string {
	cDate := commitDate.Format(vxTimeFormat)
	if UseDefaultTime {
		cDate = defaultTime
	}
	return fmt.Sprintf("%s%s%s", commitMsg, separator, cDate)
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
