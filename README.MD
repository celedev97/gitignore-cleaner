# Gitignore Cleaner

A golang project cleaner that relies on .gitignore

## What's the use of this?

One day i started giving a look at my hard drive space and i noticed i had 40GB of stuff inside the folder when i keep my repos.

Most of those repos are super old and i don't really need to have build artifacts and dependencies saved on my hard disk, they're just there wasting space.

Git has a super useful feature called [git clean](https://git-scm.com/docs/git-clean) that removes untracked files from the working directory of a git repository.

So i tought "hey why can't i just clean every repo with git clean?" but i soon realized that not every project i had was a git repository, so it was pretty much impossible to clean them all unless i would have manually git initialized them all, added a `.gitignore`, staged the files and then used `git clean`.

This litte utility comes from my necessity to clean up a large amount of projects without the need to have them being git repositories.

It find projects that have file matching a pattern and then clean them by matching their file with a `.gitignore` file.

## Getting Started

### Installation

To get this utility installed, you can simply run

```bash
go get github.com/fcdev/gitignore-cleaner
```

### Help

To have the list of the parameters just run

```
gitignore-cleaner -h
```

### Detailed flags descriptions

| flag  |example value     | action
|---    |---               |---
|-p     |repos             | Specify the folder that will be scanned for searching the projects to clean up.
|-g     |Golang.gitignore  | Specify a global `.gitignore` file used for choosing the file to delete, if not specified the program will look into the project folder for a specific `.gitignore`<br>(this could be dangerous as it could delete secrets, like android jks files that are not checked out into the repository)
|-d     |                  | Dry run mode, runs a simulation of wich file it would delete, without actually deleting anything.  
|-y     |                  | If this flag is set the program will run in batch mode deleting everything without asking for confirmation. 
|-i     | .vscode          | Specify a file that should be into a folder for that folder to be marked as a project and being cleaned up.<br> This option is used as a glob and can be used multiple times to simulate an or.
|-e     | *.jks            | This is the counterpart of `-i`, if a project should be included by `-i` but has a file specified by `-e` the **whole** project will be skipped, this is useful if you want to avoid cleaning project that might have delicate files included in the .gitignore

### Usage examples

* Cleaning all the C# projects inside the folder `repos`,
* using the rules defined in the file `CSharp.gitignore`,
* and matching as valid projects only the folders that have one `sln` or `csproj` files inside
* without asking for confirmation

```
gitignore-cleaner -p repos -g CSharp.gitignore -i *.sln -i *.csproj -y
```

## Contributing

Feel free to contact me or open a pull-request if you have some improvements in your mind, or if you found a bug.

## Safety

This program is done **for deleting stuff**, if used wrongly it can delete stuff that you don't want deleted, so be sure to use it wisely, and if you want to use the batch mode, possibly do a dry-run before that, so you can be sure you are not deleting anything important.

## Authors

* **fcdev** - *Everything for now* - [fcdev](https://github.com/fcdev)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Libraries used

* [keyboard](github.com/eiannone/keyboard) from [eiannone](github.com/eiannone) - used for the Yes/No prompt
* [go-gitignore](github.com/monochromegane/go-gitignore) from [monochromegane](github.com/monochromegane) - used for checking if a file/folder matches .gitignore
