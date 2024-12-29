package html_test

import (
	"fmt"
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

func ExampleRemoveSelections() {

	doc := []byte(`
	<!DOCTYPE html>
	<html>
		<head>
			<title>Test Title</title>
		</head>
		<body>
			<h1>Body</h1>
			<h2>Sub Body</h2>
			<p class="removeme">Remove Me</p>
			<p class="deleteme">Delete Me</p>
		</body>
	</html>`)

	newDoc, err := html.RemoveSelections(doc, []string{".removeme", ".deleteme"})
	if err != nil {
		// Handle error
	}

	fmt.Printf("%s\n", newDoc)
}
