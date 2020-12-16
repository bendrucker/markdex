package markdown

import (
	"io/ioutil"

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

	// TODO: handle heading with multiple text nodes
	var title string
	_ = ast.Walk(heading, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if text, ok := node.(*ast.Text); ok {
			title = string(text.Segment.Value(d.Source))
			return ast.WalkStop, nil
		}

		return ast.WalkContinue, nil
	})

	return title
}
