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

	expected := &Entry{
		Title: "Root",
		Slug:  "testdata/basic",
		Children: []*Entry{
			{
				Title: "Bar",
				Slug:  "testdata/basic/bar",
				Children: []*Entry{
					NewEntry("Baz", "testdata/basic/bar/baz.md"),
				},
			},
			NewEntry("Foo", "testdata/basic/foo.md"),
		},
	}

	assert.Equal(t, expected, index)
}

func TestMarkdown(t *testing.T) {
	index := &Entry{
		Title: "Root",
		Slug:  "testdata/basic",
		Children: []*Entry{
			{
				Title: "Bar",
				Slug:  "testdata/basic/bar",
				Children: []*Entry{
					NewEntry("Baz", "testdata/basic/bar/baz.md"),
				},
			},
			NewEntry("Foo", "testdata/basic/foo.md"),
		},
	}

	expected := `
* [Bar](testdata/basic/bar)
  * [Baz](testdata/basic/bar/baz.md)
* [Foo](testdata/basic/foo.md)
	`

	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(string(index.Markdown())))
}
