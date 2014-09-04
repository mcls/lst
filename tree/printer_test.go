package tree

import (
	"bytes"
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
	out := new(bytes.Buffer)
	p.Out = out
	node, _ := NewNode(BaseTestPath())
	p.Print(node, 0)
	actualLines := strings.Split(out.String(), "\n")

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
	p := &Printer{ShowAll: false, ShowOnlyDirs: false, Style: NewRegularStyle()}
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
	AssertTreeMatches(t, p, dirs, expectedTree)

	dirs = map[string]os.FileMode{
		"mydir/":               0777,
		"mydir/stuff.md":       0777,
		"mydir/thing.md":       0777,
		"mydir/more/":          0777,
		"mydir/more/stuff.md":  0777,
		"mydir/more/and.md":    0777,
		"mydir/more/things.md": 0777,
		"zzz/":                 0777,
		"zzz/naps.txt":         0777,
		"zzz/sleep.txt":        0777,
	}
	expectedTree = []string{
		".",
		"|-- mydir/",
		"|   |-- more/",
		"|   |   |-- and.md",
		"|   |   |-- stuff.md",
		"|   |   `-- things.md",
		"|   |-- stuff.md",
		"|   `-- thing.md",
		"`-- zzz/",
		"    |-- naps.txt",
		"    `-- sleep.txt",
	}
	AssertTreeMatches(t, p, dirs, expectedTree)

	p.Style = NewAnsiStyle()
	dirs = map[string]os.FileMode{
		"mydir/":               0777,
		"mydir/thing.md":       0777,
		"mydir/more/":          0777,
		"mydir/more/things.md": 0777,
		"zzz/":                 0777,
		"zzz/naps.txt":         0777,
		"zzz/sleep.txt":        0777,
	}
	expectedTree = []string{
		".",
		"├── mydir/",
		"│   ├── more/",
		"│   │   └── things.md",
		"│   └── thing.md",
		"└── zzz/",
		"    ├── naps.txt",
		"    └── sleep.txt",
	}
	AssertTreeMatches(t, p, dirs, expectedTree)
}
