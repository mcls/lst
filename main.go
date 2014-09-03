// `lst` is based on Steve Baker's `tree`
// ( http://mama.indstate.edu/users/ice/tree/ ).
//
// Built in `go` for learning purposes.
//
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mcls/lst/tree"
)

var maxLevel = flag.Uint("l", 1, "Max. level of dirs to descend (0 = infinite)")
var useAnsi = flag.Bool("A", false, "Print ANSI lines graphic indentation lines.")
var showAll = flag.Bool("a", false, "List all files")
var showOnlyDirs = flag.Bool("d", false, "Only list directories")

func main() {
	flag.Parse()
	args := flag.Args()

	// Determine root directory of tree
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}

	// Determine tree style
	if *useAnsi {
		tree.SetStyle(tree.NewAnsiStyle())
	} else {
		tree.SetStyle(tree.NewRegularStyle())
	}

	printer := &tree.Printer{
		Out:          os.Stdout,
		ShowAll:      *showAll,
		ShowOnlyDirs: *showOnlyDirs}

	// Initialize and print tree
	node, err := tree.NewNode(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	printer.Print(node, *maxLevel)
}
