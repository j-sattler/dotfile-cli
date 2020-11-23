package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
)


var (
	info = color.New(color.FgHiGreen, color.Bold)

	initCmd  = flag.NewFlagSet("init", flag.ExitOnError)
	lsCmd  = flag.NewFlagSet("ls", flag.ExitOnError)
	addCmd = flag.NewFlagSet("add", flag.ExitOnError)
	rmCmd  = flag.NewFlagSet("rm", flag.ExitOnError)
	editCmd  = flag.NewFlagSet("edit", flag.ExitOnError)
	pullCmd  = flag.NewFlagSet("pull", flag.ExitOnError)
	pushCmd  = flag.NewFlagSet("push", flag.ExitOnError)
)

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

	if len(os.Args) < 2 {
		PrintUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "ls":
		lsCmd.Parse(os.Args[2:])
		info.Println("ls command")
	case "add":
		addCmd.Parse(os.Args[2:])
		info.Println("add command")
	case "rm":
		rmCmd.Parse(os.Args[2:])
		info.Println("rm command")
	case "edit":
		editCmd.Parse(os.Args[2:])
		info.Println("edit command")
	case "init":
		initCmd.Parse(os.Args[2:])
		info.Println("init command")
	case "pull":
		pullCmd.Parse(os.Args[2:])
		info.Println("pull command")
	case "push":
		pushCmd.Parse(os.Args[2:])
		info.Println("push command")
	default:
		PrintUsage()
		os.Exit(1)
	}

	flag.Parse()

}
