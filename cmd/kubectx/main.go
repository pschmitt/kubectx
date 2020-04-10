package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	// parse command-line flags
	var op Op
	op = parseArgs(os.Args[1:])

	// TODO consider addin Run() operation to each operation type
	switch v := op.(type) {
	case HelpOp:
		printHelp(os.Stdout)
	case ListOp:
		printListContexts(os.Stdout)
	case SwitchOp:
		var newCtx string
		var err error
		if v.Target == "-" {
			newCtx, err = swapContext()
		} else {
			newCtx, err = switchContext(v.Target)
		}
		if err != nil {
			printError("failed to switch context: %v", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "Switched to context %q.\n", newCtx)
	case UnknownOp:
		printError("unsupported operation: %s", strings.Join(v.Args, " "))
		printHelp(os.Stdout)
		os.Exit(1)
	default:
		fmt.Printf("internal error: operation type %T not handled", op)
	}
}

func printError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, color.RedString("error: ")+format+"\n", args...)
}