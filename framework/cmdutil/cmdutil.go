package cmdutil

import (
	"log"
	"os"
)

// SubCommand is an implementation of a linter sub-command.
type SubCommand struct {
	// Main is command entry point.
	Main func()

	// Name is sub-command name used to execute it.
	Name string

	// Short describes command in one line of text.
	Short string

	// Examples shows one or more sub-command execution examples.
	Examples []string
}

// DispatchCommand runs sub command out of specified cmdList based on
// the first command line argument.
func DispatchCommand(cmdList []*SubCommand) {
	argv := os.Args
	if len(argv) < 2 {
		log.Printf("not enough arguments, expected sub-command name\n\n")
		printSubCommands(cmdList)
		os.Exit(1)
	}

	subIdx := 1 // [0] is program name
	sub := os.Args[subIdx]
	// Erase sub-command argument (index=1) to make it invisible for
	// sub commands themselves.
	os.Args = append(os.Args[:subIdx], os.Args[subIdx+1:]...)

	// Choose and run sub-command main.
	cmd := findSubCommand(cmdList, sub)
	if cmd == nil {
		log.Printf("unknown sub-command: %s\n\n", sub)
		printSubCommands(cmdList)
		os.Exit(1)
	}

	// The called function may exit with non-zero status.
	// No code should follow this call.
	cmd.Main()
}

// findSubCommand looks up SubCommand by its name.
// Returns nil if requested command not found.
func findSubCommand(cmdList []*SubCommand, name string) *SubCommand {
	for _, cmd := range cmdList {
		if cmd.Name == name {
			return cmd
		}
	}
	return nil
}

// printSubCommands prints cmdList info to the logger (usually stderr).
func printSubCommands(cmdList []*SubCommand) {
	log.Println("Supported sub-commands:")
	for _, cmd := range cmdList {
		log.Printf("\t%s - %s", cmd.Name, cmd.Short)
		for _, ex := range cmd.Examples {
			log.Printf("\t\t$ %s", ex)
		}
	}
}
