package cli

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
)

var (
	info = color.New(color.FgHiGreen, color.Bold)

	initCmd = flag.NewFlagSet("init", flag.ExitOnError)
	lsCmd   = flag.NewFlagSet("ls", flag.ExitOnError)
	addCmd  = flag.NewFlagSet("add", flag.ExitOnError)
	rmCmd   = flag.NewFlagSet("rm", flag.ExitOnError)
	editCmd = flag.NewFlagSet("edit", flag.ExitOnError)
	pullCmd = flag.NewFlagSet("pull", flag.ExitOnError)
	pushCmd = flag.NewFlagSet("push", flag.ExitOnError)
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

func Run(args [] string){

	if len(args) < 2 {
		PrintUsage()
		os.Exit(1)
	}

	switch args[1] {
	case "ls":
		lsCmd.Parse(args[2:])
		ls()
	case "add":
		addCmd.Parse(args[2:])
		add()
	case "rm":
		rmCmd.Parse(args[2:])
		rm()
	case "edit":
		editCmd.Parse(args[2:])
		edit()
	case "init":
		initCmd.Parse(args[2:])
		initialize()
	case "pull":
		pullCmd.Parse(args[2:])
		pull()
	case "push":
		pushCmd.Parse(args[2:])
		push()
	default:
		PrintUsage()
		os.Exit(1)
	}
}

func initialize(){
	fmt.Println("init")
}

func ls(){
	fmt.Println("ls")
}

func rm(){
	fmt.Println("rm")
}

func add(){
	fmt.Println("add")
}

func push(){
	fmt.Println("push")
}

func pull(){
	fmt.Println("pull")
}

func edit(){
	fmt.Println("edit")
}