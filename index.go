package markdex

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type Index struct {
	Entries []*Entry
}

type Entry struct {
	Title    string
	Children []*Entry
}

var markdown = goldmark.New()

func Load(dir string) (*Index, error) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// if info.IsDir() {
		// 	if file, err := os.Open()
		// }
		return nil
	})

	return nil, err
}

var readmeFilenames = []string{"readme.md", "README.md"}

func LoadReadme(dir string) (*Entry, error) {
	for _, name := range readmeFilenames {
		data, err := ioutil.ReadFile(filepath.Join(dir, name))

		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}

			return nil, err
		}

		doc := markdown.Parser().Parse(text.NewReader(data))

		var heading *ast.Heading
		for child := doc.FirstChild(); child != nil; child = child.NextSibling() {
			h, ok := child.(*ast.Heading)
			if !ok || h.Level != 1 {
				continue
			}
			heading = h
			break
		}

		if heading == nil {
			return nil, nil
		}

		// TODO: handle heading with multiple text nodes
		var title string
		err = ast.Walk(heading, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
			if text, ok := node.(*ast.Text); ok {
				title = string(text.Segment.Value(data))
				return ast.WalkStop, nil
			}

			return ast.WalkContinue, nil
		})

		if err != nil {
			return nil, err
		}

		return &Entry{Title: title}, nil
	}

	return nil, nil
}
