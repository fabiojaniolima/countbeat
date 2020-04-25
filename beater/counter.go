package beater

import (
	"fmt"
	"os"
	"path/filepath"
)

func countOnlyFiles(paths []string) int {
	var totalFiles int

	for _, file := range paths {
		stat, err := os.Stat(file)
		if err != nil {
			fmt.Println(err)
		}

		if stat.Mode().IsDir() {
			continue
		}

		totalFiles++
	}

	return totalFiles
}

func countMatches(path string, countDirectories bool) int {
	matches, err := filepath.Glob(path)

	if err != nil {
		fmt.Println(err)
	}

	if countDirectories {
		return len(matches)
	}

	return countOnlyFiles(matches)
}
