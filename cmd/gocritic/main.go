package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-critic/go-critic/cmd/criticize"
	"github.com/go-critic/go-critic/cmd/lintwalk"
)

var version = "0.3.3"

// subCommand is an implementation of a gocritic sub-command.
type subCommand struct {
	// main is command entry point.
	main func()

	// name is sub-command name used to execute it.
	name string

	// short describes command in one line of text.
	short string
}

// subCommands describes all supported sub-commands as well
// as their metadata required to run them and print useful help messages.
var subCommands = []*subCommand{
	{
		main:  criticize.Main,
		name:  "check-package",
		short: "run gocritic over specified package list",
	},
	{
		main:  lintwalk.Main,
		name:  "check-project",
		short: "run gocritic over specified source tree, recursively",
	},
	{
		main:  printVersion,
		name:  "version",
		short: "print gocritic version",
	},
}

func main() {
	log.SetFlags(0)

	argv := os.Args
	if len(argv) < 2 {
		terminate("not enough arguments, expected sub-command name", printUsage)
	}

	subIdx := 1 // [0] is program name
	sub := os.Args[subIdx]
	// Erase sub-command argument (index=1) to make it invisible for
	// sub commands themselves.
	os.Args = append(os.Args[:subIdx], os.Args[subIdx+1:]...)

	// Choose and run sub-command main.
	cmd := findSubCommand(sub)
	if cmd == nil {
		terminate("unknown sub-command: "+sub, printSupportedSubs)
	}
	log.SetPrefix(sub + ": ")
	// The called function may exit with non-zero status.
	// No code should follow this call.
	cmd.main()
}

func printVersion() {
	log.Println(version)
}

// findSubCommand looks up subCommand by its name.
// Returns nil if requested command not found.
func findSubCommand(name string) *subCommand {
	for _, cmd := range subCommands {
		if cmd.name == name {
			return cmd
		}
	}
	return nil
}

// terminate prints error specified by reason, runs optional printHelp
// function and then exists with non-zero status.
func terminate(reason string, printHelp func()) {
	stderrPrintf("error: %s\n", reason)
	if printHelp != nil {
		stderrPrintf("\n")
		printHelp()
	}
	os.Exit(1)
}

func printUsage() {
	// TODO: implement me. For now, print supported commands.
	printSupportedSubs()
}

func printSupportedSubs() {
	stderrPrintf("Supported sub-commands:\n")
	for _, cmd := range subCommands {
		stderrPrintf("\t%s - %s\n", cmd.name, cmd.short)
	}
}

// stderrPrintf writes formatted message to stderr and checks for error
// making "not annoying at all" linters happy.
func stderrPrintf(format string, args ...interface{}) {
	_, err := fmt.Fprintf(os.Stderr, format, args...)
	if err != nil {
		panic(fmt.Sprintf("stderr write error: %v", err))
	}
}
