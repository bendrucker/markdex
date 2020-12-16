package markdex

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/bendrucker/markdex/internal/markdown"
)

type Index struct {
	Entries []*Entry
}

type Entry struct {
	Title    string
	Children []*Entry
}

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
		doc, err := markdown.Load(filepath.Join(dir, name))

		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}

			return nil, err
		}

		return &Entry{Title: doc.Title()}, nil
	}

	return nil, nil
}
