package markdex

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	index, err := Load("testdata/basic")
	if err != nil {
		t.Fatal(err)
	}

	expected := newParentNode("Root", "testdata/basic", []*Node{
		newParentNode("Bar", "bar", []*Node{
			newNode("Baz", "baz.md"),
		}),
		newNode("Foo", "foo.md"),
	})

	assert.Equal(t, expected, index)
}

func TestLoad_noReadme(t *testing.T) {
	index, err := Load("testdata/no-readme")
	if err != nil {
		t.Fatal(err)
	}

	expected := newParentNode("Root", "testdata/no-readme", []*Node{
		newParentNode("foo", "foo", []*Node{
			newNode("Bar", "bar.md"),
		}),
	})

	assert.Equal(t, expected, index)
}

func TestLoadReadme_lowercase(t *testing.T) {
	readme, err := LoadReadme("testdata/lowercase")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Lowercase Filename", readme.Title)
}

func TestMarkdown(t *testing.T) {
	index := newParentNode("Root", "testdata/basic", []*Node{
		newParentNode("Bar", "bar", []*Node{
			newNode("Baz", "baz.md"),
		}),
		newNode("Foo", "foo.md"),
	})

	expected := `
* [Bar](testdata/basic/bar)
  * [Baz](testdata/basic/bar/baz.md)
* [Foo](testdata/basic/foo.md)
	`

	actual := string(Markdown(index))

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(actual))
}
