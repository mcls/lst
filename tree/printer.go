package tree

import (
	"io"
	"os"
	"path/filepath"
)

// Printer contains the options for printing the tree
type Printer struct {
	ShowAll      bool
	ShowOnlyDirs bool
	Out          io.Writer
	Style        *Style
}

// Print outputs the tree
func (printer *Printer) Print(node *Node, maxLevel uint) {
	if node.Level == 0 {
		printer.printRelativePath(node.AbsPath)
	}
	nodes := node.Children()
	if maxLevel > 0 && node.Level >= maxLevel {
		return
	}

	for _, node := range nodes {
		if !printer.shouldPrintNode(node) {
			continue
		} else {
			printer.printLine(node.Line(printer.Style))
			if node.FileInfo.IsDir() {
				printer.Print(node, maxLevel)
			}
		}
	}
}

func (printer *Printer) printLine(txt string) {
	printer.Out.Write([]byte(txt + "\n"))
}

func (printer *Printer) printRelativePath(absPath string) {
	base, _ := os.Getwd()
	path, _ := filepath.Rel(base, absPath)
	printer.printLine(path)
}

func (printer *Printer) shouldPrintNode(node *Node) bool {
	if !printer.ShowAll && node.IsDotfile() {
		return false
	}
	if printer.ShowOnlyDirs && !node.FileInfo.IsDir() {
		return false
	}
	return true
}
