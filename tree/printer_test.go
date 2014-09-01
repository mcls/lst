package tree

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestPrintingDir(t *testing.T) {
	f, _ := ioutil.TempFile("", "testing")
	p := &Printer{
		ShowAll:      false,
		ShowOnlyDirs: false,
		Out:          f}
	node, _ := NewNode(".")
	p.Print(node, 0)

	content, _ := ioutil.ReadFile(f.Name())
	actual := string(content)

	lines := []string{
		".",
		"|-- node.go",
		"|-- printer.go",
		"|-- printer_test.go",
		"`-- style.go"}

	actualLines := strings.Split(actual, "\n")
	for i, line := range lines {
		if i < len(actualLines) {
			actualLine := actualLines[i]
			if line != actualLine {
				t.Errorf("Expected actualLines[%d] to == '%s' was '%s'", i, line, actualLine)
			}
		} else {
			t.Errorf("Expected actualLines[%d] to == '%s' was out of bounds", i, line)
		}
	}
}
