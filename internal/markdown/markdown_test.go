package markdown

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestDocumentTitle(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		Name     string
		Source   string
		Expected string
	}{
		{
			Name:     "standard",
			Expected: "Standard Title",
		},
		{
			Name:     "underscore",
			Expected: "Underscored Title",
		},
		{
			Name:     "empty",
			Expected: "",
		},
		{
			Name:     "h2",
			Expected: "",
		},
		{
			Name:     "formatted",
			Expected: "Italics and a Link",
		},
		{
			Name:     "image",
			Expected: "Title With Inline Image",
		},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			doc, err := Load(filepath.Join("testdata", fmt.Sprintf("%s.md", tc.Name)))
			if err != nil {
				t.Fatalf("failed to load document: %s", err)
			}

			if got := doc.Title(); got != tc.Expected {
				t.Errorf("Title() = %s, expected %s", got, tc.Expected)
				doc.Dump()
			}
		})
	}
}
