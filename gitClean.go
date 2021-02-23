package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/monochromegane/go-gitignore"
)

// GitClean cleans every subproject in a folder from the files not in gitignore
func GitClean(path string, gitignorePath string, includeGlobs []string, excludedGlobs []string, dryRun bool, automatic bool) {
	projects := getProjects(path, includeGlobs, excludedGlobs)

	for _, project := range projects {
		println("PROCESSING: ", project)

		toUseGitIgnore := filepath.Join(project, ".gitignore")
		if gitignorePath != "" {
			toUseGitIgnore = gitignorePath
		}

		gitIgnoreHelper, err := gitignore.NewGitIgnore(toUseGitIgnore, project)
		if err != nil {
			println("ERROR CAN'T CREATE GITIGNORE HELPER", err.Error())
			continue
		}

		filepath.Walk(project, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			if gitIgnoreHelper.Match(path, info.IsDir()) {
				fmt.Println("REMOVING:", path, " ")
				if !dryRun && (automatic || yesno("Confirm ", true)) {
					err = os.RemoveAll(path)
					if err != nil {
						println("CANNOT REMOVE: ", path, ":", err.Error())
					} else {
						fmt.Println("REMOVED")
					}
				}

			}
			return nil
		})
		println("DONE: ", project, "\n")
	}

}
