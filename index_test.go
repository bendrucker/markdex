package markdex

import (
	"testing"
)

func TestLoadReadme(t *testing.T) {
	entry, err := LoadReadme("testdata/readme")
	if err != nil {
		t.Fatal(err)
	}

	if entry.Title != "Readme Title" {
		t.Fatalf("unexpected title: %s", entry.Title)
	}
}
