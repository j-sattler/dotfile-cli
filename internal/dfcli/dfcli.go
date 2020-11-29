package dfcli

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"os"
)

type Dotfile struct {
	confPath string
}

const (
	dotfileDirName = ".dotfile"
)

var (
	info = color.New(color.FgHiGreen, color.Bold)
)

func printUsage() {
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

func Run() {

	// Check if subcommand is present
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Find and execute subcommand
	switch os.Args[1] {
	case "ls":
		ls()
	case "add":
		add()
	case "rm":
		rm()
	case "edit":
		edit()
	case "init":
		initialize()
	case "pull":
		pull()
	case "push":
		push()
	case "help":
		help()
	default:
		printUsage()
		os.Exit(1)
	}
}

func initialize() {
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	initCmd.Parse(os.Args[2:])

	dotfile := Dotfile{}
	dotfile.determineConfDir()

	err := os.MkdirAll(dotfile.confPath, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create config file directory: %s\n", err)
		os.Exit(1)
	}

	_ , err = git.PlainInit(dotfile.confPath, false)

	if err != nil {
		fmt.Printf("Failed to initialize git repository: %s\n", err)
	}


	_, _ = info.Printf("Initialized empty dotfile project in %s\n", dotfile.confPath)

}

// determine the config directory based on user's OS config or home directory
func (d *Dotfile) determineConfDir() {
	dir, err := os.UserConfigDir()

	// check if user config dir exists
	if err != nil {
		// could not determine user config dir
		// try user home dir as fallback
		dir , err = os.UserHomeDir()
		if err != nil {
			// could neither determine home nor config dir
			fmt.Printf("Could not determine user home or config directory: %s\n", err)
			os.Exit(1)
		}
	}
	// if place dotfile config
	d.confPath = fmt.Sprintf("%s/%s", dir, dotfileDirName)
}

func ls() {
	lsCmd := flag.NewFlagSet("ls", flag.ExitOnError)
	lsCmd.Parse(os.Args[2:])
	fmt.Println(lsCmd.Args())
}

func rm() {
	rmCmd := flag.NewFlagSet("rm", flag.ExitOnError)
	rmCmd.Parse(os.Args[2:])
	fmt.Println(rmCmd.Args())
}

func add() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	fmt.Println(addCmd.Args())
}

func push() {
	pushCmd := flag.NewFlagSet("push", flag.ExitOnError)
	fmt.Println(pushCmd.Args())
}

func pull() {
	pullCmd := flag.NewFlagSet("pull", flag.ExitOnError)
	fmt.Println(pullCmd.Args())
}

func edit() {
	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
	fmt.Println(editCmd.Args())
}

func help() {
	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)
	fmt.Println(helpCmd.Args())
	printUsage()
	os.Exit(1)
}

