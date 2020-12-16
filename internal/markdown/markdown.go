package markdown

import (
	"io/ioutil"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

var md = goldmark.New()

// Load loads and parses a markdown document, returning an object with both the parsed AST and original source
func Load(path string) (*Document, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return Parse(data), nil
}

// Parse parses markdown data into a Document
func Parse(data []byte) *Document {
	return &Document{
		AST:    md.Parser().Parse(text.NewReader(data)),
		Source: data,
	}
}

// Document represents a Markdown document
type Document struct {
	AST    ast.Node
	Source []byte
}

// Dump prints the Document's AST to stdout for debugging
func (d *Document) Dump() {
	d.AST.Dump(d.Source, 0)
}

// Title returns the title of a markdown document from the first <h1>
func (d *Document) Title() string {
	var heading *ast.Heading
	for child := d.AST.FirstChild(); child != nil; child = child.NextSibling() {
		h, ok := child.(*ast.Heading)
		if !ok || h.Level != 1 {
			continue
		}
		heading = h
		break
	}

	if heading == nil {
		return ""
	}

	parts := []string{}
	_ = ast.Walk(heading, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if text, ok := node.(*ast.Text); entering && ok {
			parts = append(parts, string(text.Segment.Value(d.Source)))
		}

		if _, ok := node.(*ast.Image); ok {
			return ast.WalkSkipChildren, nil
		}

		return ast.WalkContinue, nil
	})

	return strings.TrimSpace(strings.Join(parts, ""))
}
