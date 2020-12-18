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

func TestLoad_noReadme(t *testing.T) {
	index, err := Load("testdata/no-readme")
	if err != nil {
		t.Fatal(err)
	}

	expected := &Entry{
		Title: "Root",
		Slug:  "testdata/no-readme",
		Children: []*Entry{
			{
				Title: "foo",
				Slug:  "testdata/no-readme/foo",
				Children: []*Entry{
					NewEntry("Bar", "testdata/no-readme/foo/bar.md"),
				},
			},
		},
	}

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
