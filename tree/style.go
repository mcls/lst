package tree

// Style determines how the tree is drawn
type Style struct {
	middle          string
	last            string
	indentNested    string
	indentNotNested string
}

// NewRegularStyle returns a style that will draw the tree using regular
// characters
func NewRegularStyle() *Style {
	return &Style{
		middle:          "|-- ",
		last:            "`-- ",
		indentNested:    "|   ",
		indentNotNested: "    "}
}

// NewAnsiStyle returns a style that will draw the tree using ANSI characters
func NewAnsiStyle() *Style {
	return &Style{
		middle:          "├── ",
		last:            "└── ",
		indentNested:    "│   ",
		indentNotNested: "    "}
}
