# dotfile-cli
dotfile is a tool for managing dotfiles

## Features
- manage dotfiles in a central place
- sync with git

## Commands

```
Usage:
	dotfile <command> [arguments]

The commands are:

    init    initialize new dotfile project
    add     add dotfile to project
    rm	    remove dotfile from project
    ls	    list all dotfiles in project
    push    push dotfiles to git repository
    pull    pull dotfiles from git repository
    edit    edit dotfile
    remote  configure git remote repo
```

## Install from source

From the root of this project run:

```bash
go install cmd/dotfile/dotfile.go
``` 

## To Do

- [x] init command
- [ ] add command
- [ ] rm command
- [ ] ls command
- [ ] push command
- [ ] pull command
- [ ] edit command
- [ ] remote command

## Third Party Libraries

* [**go-git/go-git**](https://github.com/go-git/go-git)
* [**fatih/color**](https://github.com/fatih/color)