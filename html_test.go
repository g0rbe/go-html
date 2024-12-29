package html_test

import (
	"testing"

	"git.gorbe.io/go/html"
)

func TestRemoveSelections(t *testing.T) {

	newHtml, err := html.RemoveSelections([]byte(TestDocument), []string{".removeme", ".clearme"})
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	t.Logf("\n%s\n", newHtml)
}
