// `lst` is based on Steve Baker's `tree`
// ( http://mama.indstate.edu/users/ice/tree/ ).
//
// Built in `go` for learning purposes.
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var maxLevel = flag.Uint("level", 0, "Max. level of dirs to descend (0 = infinite)")
var useAnsi = flag.Bool("A", false, "Print ANSI lines graphic indentation lines.")
var showAll = flag.Bool("a", false, "List all files")
var showOnlyDirs = flag.Bool("d", false, "Only list directories")

// LineGraphics determine the style of indentation
type LineGraphics struct {
	middle          string
	last            string
	indentNested    string
	indentNotNested string
}

var lineGraphics LineGraphics

func main() {
	flag.Parse()
	args := flag.Args()

	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}

	if *useAnsi {
		lineGraphics = newAnsiLineGraphics()
	} else {
		lineGraphics = newRegularLineGraphics()
	}

	node := initTreeNode(dir)
	fmt.Println(dir)
	printTree(node, *maxLevel)
}

func newRegularLineGraphics() LineGraphics {
	return LineGraphics{
		middle:          "|-- ",
		last:            "`-- ",
		indentNested:    "|   ",
		indentNotNested: "    "}
}

func newAnsiLineGraphics() LineGraphics {
	return LineGraphics{
		middle:          "├── ",
		last:            "└── ",
		indentNested:    "│   ",
		indentNotNested: "    "}
}

func printTree(node *TreeNode, maxLevel uint) {
	nodes := node.Children()
	if maxLevel > 0 && node.level >= maxLevel {
		return
	}

	for _, node := range nodes {
		if !*showAll && node.IsDotfile() {
			continue
		}
		if *showOnlyDirs && !node.fileInfo.IsDir() {
			continue
		}
		fmt.Println(node.Line())
		if node.fileInfo.IsDir() {
			printTree(&node, maxLevel)
		}
	}
}

func initTreeNode(dir string) *TreeNode {
	fileInfo, _ := os.Stat(dir)
	absPath, _ := filepath.Abs(dir)
	return &TreeNode{
		fileInfo:    fileInfo,
		indentation: "",
		absPath:     absPath,
		isLast:      false,
		level:       0}
}

// TreeNode is a wrapper for FileInfo with extra information such as
// indentation and level
type TreeNode struct {
	fileInfo    os.FileInfo
	level       uint
	isLast      bool
	indentation string
	absPath     string
}

// Name of the associated file, suffixes separator (e.g. "/") when directory
func (node *TreeNode) Name() string {
	if node.fileInfo.IsDir() {
		return node.fileInfo.Name() + string(os.PathSeparator)
	}
	return node.fileInfo.Name()
}

// Line returns the name of the file with indentation
func (node *TreeNode) Line() string {
	branch := lineGraphics.middle
	if node.isLast {
		branch = lineGraphics.last
	}
	return node.indentation + branch + node.Name()
}

// IsDotfile returns true when the file name starts with a '.'
func (node *TreeNode) IsDotfile() bool {
	return strings.Index(node.fileInfo.Name(), ".") == 0
}

func (node *TreeNode) childIndentation() (indentation string) {
	if node.level > 0 {
		var indent string
		if node.isLast {
			indent = lineGraphics.indentNotNested
		} else {
			indent = lineGraphics.indentNested
		}
		indentation = node.indentation + indent
	}
	return
}

// Children returns all files nested in a directory as []TreeNode
func (node *TreeNode) Children() []TreeNode {
	entries, err := ioutil.ReadDir(node.absPath)
	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Println("PATH: ", node.absPath)
		os.Exit(1)
	}
	// Wrap in TreeNode struct
	nodes := make([]TreeNode, len(entries))
	for i, entry := range entries {
		nodes[i] = TreeNode{
			fileInfo:    entry,
			level:       node.level + 1,
			isLast:      len(entries)-1 == i,
			absPath:     filepath.Join(node.absPath, entry.Name()),
			indentation: node.childIndentation()}
	}
	return nodes
}
