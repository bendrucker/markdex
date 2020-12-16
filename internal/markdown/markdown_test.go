package markdown

import (
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
			Source:   "# Title",
			Expected: "Title",
		},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			doc := Parse([]byte(tc.Source))
			if got := doc.Title(); got != tc.Expected {
				t.Errorf("Title() = %s, expected %s", got, tc.Expected)
			}
		})
	}
}
