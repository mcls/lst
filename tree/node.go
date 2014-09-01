package tree

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Node is a wrapper for FileInfo with extra information such as
// indentation and level
type Node struct {
	FileInfo    os.FileInfo
	Level       uint
	IsLast      bool
	Indentation string
	AbsPath     string
}

// NewNode creates a new Node for the given directory
func NewNode(dir string) (*Node, error) {
	fileInfo, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	absPath, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	return &Node{
		FileInfo:    fileInfo,
		Indentation: "",
		AbsPath:     absPath,
		IsLast:      false,
		Level:       0}, nil
}

// Name of the associated file, suffixes separator (e.g. "/") when directory
func (node *Node) Name() string {
	if node.FileInfo.IsDir() {
		return node.FileInfo.Name() + string(os.PathSeparator)
	}
	return node.FileInfo.Name()
}

// Line returns the name of the file with indentation
func (node *Node) Line() string {
	branch := style.middle
	if node.IsLast {
		branch = style.last
	}
	return node.Indentation + branch + node.Name()
}

// IsDotfile returns true when the file name starts with a '.'
func (node *Node) IsDotfile() bool {
	return strings.Index(node.FileInfo.Name(), ".") == 0
}

func (node *Node) childIndentation() (indentation string) {
	if node.Level > 0 {
		var indent string
		if node.IsLast {
			indent = style.indentNotNested
		} else {
			indent = style.indentNested
		}
		indentation = node.Indentation + indent
	}
	return
}

// Children returns all files nested in a directory as []Node
func (node *Node) Children() []Node {
	entries, err := ioutil.ReadDir(node.AbsPath)
	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Println("PATH: ", node.AbsPath)
		os.Exit(1)
	}
	// Wrap in Node struct
	nodes := make([]Node, len(entries))
	for i, entry := range entries {
		nodes[i] = Node{
			FileInfo:    entry,
			Level:       node.Level + 1,
			IsLast:      len(entries)-1 == i,
			AbsPath:     filepath.Join(node.AbsPath, entry.Name()),
			Indentation: node.childIndentation()}
	}
	return nodes
}
