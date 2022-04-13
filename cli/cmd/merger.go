package cmd

import (
	"bufio"
	"io"
)

func merger(oldSrc, newSrc io.Reader) []string {
	oldLineMap := getLineMap(oldSrc)
	newLineMap := getLineMap(newSrc)
	diffLineMap := calcDiffLineMap(oldLineMap, newLineMap)

	var mergeContent []string

	for line := 1; line <= len(oldLineMap)+len(diffLineMap); line++ {
		diff, isDifference := diffLineMap[line]
		if isDifference {
			mergeContent = append(mergeContent, diff)
		} else {
			old, ok := oldLineMap[line]
			if ok {
				mergeContent = append(mergeContent, old)
			}
		}
	}

	return mergeContent
}

func calcDiffLineMap(oldLineMap, newLineMap map[int]string) map[int]string {
	differenceMap := make(map[int]string)
	for lineNo, newSentence := range newLineMap {
		val, ok := oldLineMap[lineNo]
		if !ok { // it means new line added
			differenceMap[lineNo] = newSentence
			continue
		}

		if newSentence != val { // it means new line updated
			differenceMap[lineNo] = newSentence
		}
	}
	return differenceMap
}

func getLineMap(reader io.Reader) map[int]string {
	scanner := bufio.NewScanner(reader)
	lineMap := make(map[int]string, 0)
	lineCount := 1
	for scanner.Scan() {
		lineMap[lineCount] = scanner.Text()
		lineCount++
	}
	return lineMap
}
