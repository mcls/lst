package tree

import (
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"testing"
)

func SetupDirStructure(dirs map[string]os.FileMode) {
	base := BaseTestPath()
	// Reset test directory
	os.RemoveAll(base)
	os.MkdirAll(base, 0777)

	keys := make([]string, 0, len(dirs))
	for dir := range dirs {
		keys = append(keys, dir)
	}
	sort.Strings(keys)

	for _, name := range keys {
		path := base + "/" + name
		mode := dirs[name]
		if strings.HasSuffix(path, "/") {
			os.MkdirAll(path, mode)
		} else {
			os.Create(path)
			os.Chmod(path, mode)
		}
	}
}

func AssertTreeMatches(t *testing.T, p *Printer, dirs map[string]os.FileMode, expectedLines []string) {
	SetupDirStructure(dirs)

	// Print tree
	out, _ := ioutil.TempFile("", "testing")
	p.Out = out
	node, _ := NewNode(BaseTestPath())
	p.Print(node, 0)
	content, _ := ioutil.ReadFile(out.Name())
	actual := string(content)
	actualLines := strings.Split(actual, "\n")

	// Compare to expected output
	for i, line := range expectedLines {
		if i == 0 {
			continue
		}
		if i < len(actualLines) {
			actualLine := actualLines[i]
			if line != actualLine {
				msg := strings.Join([]string{
					"Expected actualLines[%d]",
					"to be  : %q",
					"but was: %q",
				}, "\n")
				t.Errorf(msg, i, line, actualLine)
				return
			}
		} else {
			t.Errorf("Expected actualLines[%d] to be %q, but was out of bounds", i, line)
			return
		}
	}
}

func BaseTestPath() string {
	return os.TempDir() + "tmp/tests"
}

func TestPrintingDir(t *testing.T) {
	dirs := map[string]os.FileMode{
		"mydir/":         0777,
		"mydir/stuff.md": 0777,
		"mydir/thing.md": 0777,
		"hello.txt":      0777,
		"world.txt":      0777,
	}
	expectedTree := []string{
		".",
		"|-- hello.txt",
		"|-- mydir/",
		"|   |-- stuff.md",
		"|   `-- thing.md",
		"`-- world.txt",
	}
	p := &Printer{ShowAll: false, ShowOnlyDirs: false}
	AssertTreeMatches(t, p, dirs, expectedTree)
}
