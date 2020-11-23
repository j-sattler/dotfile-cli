package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
)

var info = color.New(color.FgHiGreen, color.Bold)

func PrintUsage() {
	fmt.Print(
		"dotfile is a tool for managing dotfiles.\n\n" +
			"Usage:\n" +
			"\tdotfile <command> [arguments]\n\n" +
			"The commands are:\n\n" +
			"\tinit\tinitialize new dotfile project\n" +
			"\tadd\tadd dotfile to project\n" +
			"\trm\tremove dotfile from project\n" +
			"\tls\tlist all dotfiles in project\n" +
			"\tpush\tpush dotfiles to git repository\n" +
			"\tpull\tpull dotfiles from git repository\n" +
			"\tedit\tedit dotfile\n\n")
}

func main() {
	listSet := flag.NewFlagSet("ls", flag.ExitOnError)
	addSet := flag.NewFlagSet("add", flag.ExitOnError)
	removeSet := flag.NewFlagSet("rm", flag.ExitOnError)

	if len(os.Args) < 2 {
		PrintUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "ls":
		listSet.Parse(os.Args[2:])
		info.Println("ls command")
	case "add":
		addSet.Parse(os.Args[2:])
		info.Println("add command")
	case "rm":
		removeSet.Parse(os.Args[2:])
		info.Println("rm command")
	case "edit":
		info.Println("edit command")
	case "init":
		info.Println("init command")
	case "pull":
		info.Println("pull command")
	case "push":
		info.Println("push command")
	default:
		PrintUsage()
		os.Exit(1)
	}

	flag.Parse()

}
