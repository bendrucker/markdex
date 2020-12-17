package markdex

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bendrucker/markdex/internal/markdown"
)

type Index struct {
	Entries []*Entry
}

type Entry struct {
	Title    string
	Slug     string
	Children []*Entry
}

func NewEntry(title, slug string) *Entry {
	return &Entry{
		Title:    title,
		Slug:     slug,
		Children: []*Entry{},
	}
}

func Load(dir string) (*Entry, error) {
	root := &Entry{
		Title:    "Root",
		Slug:     dir,
		Children: []*Entry{},
	}

	load(dir, root)
	return root, nil

	// entry := root
	// err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if path == dir {
	// 		return nil
	// 	}

	// 	if info.IsDir() {
	// 		readme, err := LoadReadme(path)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		if readme == nil {
	// 			readme = NewEntry(filepath.Base(path), path)
	// 		}

	// 		entry.Children = append(entry.Children, readme)
	// 		entry = readme
	// 		return nil
	// 	}

	// 	for _, name := range readmeFilenames {
	// 		if name == filepath.Base(path) {
	// 			return nil
	// 		}
	// 	}

	// 	doc, err := markdown.Load(path)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	entry.Children = append(entry.Children, NewEntry(doc.Title(), path))

	// 	return nil
	// })

	// return root, err
}

func load(path string, entry *Entry) error {
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

Infos:
	for _, info := range infos {
		itemPath := filepath.Join(path, info.Name())

		if info.IsDir() {
			readme, err := LoadReadme(itemPath)
			if err != nil {
				return err
			}

			if readme == nil {
				readme = NewEntry(filepath.Base(itemPath), itemPath)
			}

			entry.Children = append(entry.Children, readme)
			load(itemPath, readme)
		} else {
			for _, name := range readmeFilenames {
				if name == filepath.Base(itemPath) {
					continue Infos
				}
			}

			doc, err := markdown.Load(itemPath)
			if err != nil {
				return err
			}

			entry.Children = append(entry.Children, NewEntry(doc.Title(), itemPath))
		}
	}

	return nil
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

		return NewEntry(doc.Title(), dir), nil
	}

	return nil, nil
}
