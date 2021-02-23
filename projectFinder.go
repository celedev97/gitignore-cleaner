package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func isProject(path string, includeGlobs []string, excludedGlobs []string) bool {
	valid := false
	for _, rule := range includeGlobs {
		matches, err := filepath.Glob(filepath.Join(path, rule))
		if err == nil && len(matches) > 0 {
			valid = true
			fmt.Println(path, " included because of: ", rule)
			break
		}
	}

	for _, rule := range excludedGlobs {
		matches, err := filepath.Glob(filepath.Join(path, rule))
		if err == nil && len(matches) > 0 {
			valid = false
			fmt.Println(path, " EXCLUDED because of: ", rule)
			break
		}
	}

	return valid
}

func getProjects(path string, includeGlobs []string, excludedGlobs []string) []string {
	output := []string{}

	//check if this dir is a projects
	if isProject(path, includeGlobs, excludedGlobs) {
		output = append(output, path)
		return output
	}

	//if it's not a projects recursively scan subdirectories
	files, err := os.ReadDir(path)
	if err != nil {
		return output
	}
	for _, file := range files {
		if file.IsDir() {
			output = append(output, getProjects(filepath.Join(path, file.Name()), includeGlobs, excludedGlobs)...)
		}
	}

	return output
}
