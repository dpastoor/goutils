package goutils

import (
	"strings"
)

// ListFilesByExt returns an array of model files given an array of files of any type
func ListFilesByExt(files []string, modExt string) []string {
	var matchFiles []string
	for _, file := range files {
		if strings.HasSuffix(file, modExt) {
			matchFiles = append(matchFiles, file)
		}
	}
	return matchFiles
}
