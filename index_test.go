package markdex

import (
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
