package markdex

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bendrucker/markdex/internal/markdown"
)

func Load(dir string) (*Node, error) {
	root := newNode("Root", dir)

	load(dir, root)
	return root, nil
}

func load(path string, node *Node) error {
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
				readme = newNode(info.Name(), info.Name())
			}

			node.AppendChild(readme)
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

			node.AppendChild(newNode(doc.Title(), info.Name()))
		}
	}

	return nil
}

var readmeFilenames = []string{"readme.md", "README.md"}

func LoadReadme(dir string) (*Node, error) {
	for _, name := range readmeFilenames {
		doc, err := markdown.Load(filepath.Join(dir, name))

		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}

			return nil, err
		}

		return newNode(doc.Title(), filepath.Base(dir)), nil
	}

	return nil, nil
}

func Markdown(node *Node) []byte {
	buffer := &bytes.Buffer{}
	for _, child := range node.Children {
		toBullets(buffer, child, 0)
	}
	return buffer.Bytes()
}

func toBullets(buffer *bytes.Buffer, n *Node, level int) {
	_, _ = buffer.WriteString(fmt.Sprintf(
		"%s* [%s](%s)\n",
		strings.Repeat("  ", level),
		n.Title,
		n.Path(),
	))

	for _, child := range n.Children {
		toBullets(buffer, child, level+1)
	}
}
