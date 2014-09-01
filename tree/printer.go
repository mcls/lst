package tree

import (
	"io"
	"os"
	"path/filepath"
)

type Printer struct {
	ShowAll      bool
	ShowOnlyDirs bool
	Out          io.Writer
}

func (printer *Printer) Print(node *Node, maxLevel uint) {
	if node.Level == 0 {
		base, _ := os.Getwd()
		path, _ := filepath.Rel(base, node.AbsPath)
		printer.Out.Write([]byte(path + "\n"))
	}
	nodes := node.Children()
	if maxLevel > 0 && node.Level >= maxLevel {
		return
	}

	for _, node := range nodes {
		if !printer.ShowAll && node.IsDotfile() {
			continue
		}

		if printer.ShowOnlyDirs && !node.FileInfo.IsDir() {
			continue
		}

		printer.Out.Write([]byte(node.Line() + "\n"))

		if node.FileInfo.IsDir() {
			printer.Print(&node, maxLevel)
		}
	}
}
