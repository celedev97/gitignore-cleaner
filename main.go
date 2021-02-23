package main

import (
	"flag"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}
func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// go-gitdestroy   -e="secrets" -e "asdf" -g="./Java.gitignore" -y
func main() {
	//parsing command line arguments
	var includedGlobs arrayFlags
	var excludedGlobs arrayFlags
	var automatic *bool
	var dryRun *bool
	var gitIgnorePath *string
	var path *string

	path = flag.String("p", ".", "Path to be processed")
	gitIgnorePath = flag.String("g", "", "Specify the path of the gitignore file to use for cleaning ignored files.")
	automatic = flag.Bool("y", false, "Automatic mode, confirms all the projects for processing.")
	dryRun = flag.Bool("d", false, "Dry run mode, doesn't really delete files.")
	flag.Var(&includedGlobs, "i", "Included globs for determining if a directory is a project to be processed.")
	flag.Var(&excludedGlobs, "e", "Excluded globs for determining if a directory is a project to be processed.")

	flag.Parse()

	//printing command line arguments
	println("SETTINGS...")
	println("PATH:", *path)
	println("GITIGNORE:", *gitIgnorePath)
	println("AUTOMATIC:", *automatic)
	println("DRYRUN:", *dryRun)
	for _, v := range includedGlobs {
		println("INCLUDED:", v)
	}
	for _, v := range excludedGlobs {
		println("EXCLUDED:", v)
	}

	//running program
	GitClean(*path, *gitIgnorePath, includedGlobs, excludedGlobs, *dryRun, *automatic)

}
