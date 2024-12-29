package html_test

import (
	"bytes"
	"testing"

	"git.gorbe.io/go/html"
)

var TestDocument = `<!DOCTYPE html>
<html>
	<head>
		<title>Test Title</title>
		<meta name="description" content="Test Description">
		<link rel="canonical" href="https://example.com/">
		
		<meta property="og:url" content="https://example.com/">
		<meta property="og:site_name" content="Example">
		<meta property="og:title" content="OG Title">
		<meta property="og:description" content="OG Description">
		<meta property="og:locale" content="en"><meta property="og:type" content="website">
	</head>
	<body>
		<h1>Test Body</h1>
		<h2>Test Sub Body</h2>
		<p class="removeme">Test Class</p>
		
		<div class="deleteme">
			<p class="clearme">Test Class</p>
			<p class="clearme">Test Class</p>
		</div>
	</body>
</html>
`

// func TestDocumentParse(t *testing.T) {

// 	doc, err := html.ParseDocument([]byte(TestDocument))
// 	if err != nil {
// 		t.Fatalf("Failed to parse document: %s\n", err)
// 	}

// 	headBytes, err := doc.HeadBytes()
// 	if err != nil {
// 		t.Fatalf("Failed to get head bytes: %s\n", err)
// 	}

// 	t.Logf("%s\n", headBytes)

// 	bodyBytes, err := doc.BodyBytes()
// 	if err != nil {
// 		t.Fatalf("Failed to get body bytes: %s\n", err)
// 	}

// 	t.Logf("%s\n", bodyBytes)
// }

func TestDocumentHTML(t *testing.T) {

	d, err := html.ParseDocument([]byte(TestDocument))
	if err != nil {
		t.Fatalf("Failed to parse document: %s\n", err)
	}

	v := d.OuterHTML("head meta[property*=\"og:\"][content]")
	t.Logf("\n%s\n", v)
}

func TestDocumentTitle(t *testing.T) {

	d, err := html.ParseDocument([]byte(TestDocument))
	if err != nil {
		t.Fatalf("Failed to parse document: %s\n", err)
	}

	title := d.Title()
	if !bytes.Equal(title, []byte("Test Title")) {
		t.Fatalf("Invalid title: %s\n", title)
	}
}

func TestDocumentDescription(t *testing.T) {

	d, err := html.ParseDocument([]byte(TestDocument))
	if err != nil {
		t.Fatalf("Failed to parse document: %s\n", err)
	}

	desc := d.Description()
	if !bytes.Equal(desc, []byte("Test Description")) {
		t.Fatalf("Invalid description: %s\n", desc)
	}
}

func TestDocumentCanonical(t *testing.T) {

	d, err := html.ParseDocument([]byte(TestDocument))
	if err != nil {
		t.Fatalf("Failed to parse document: %s\n", err)
	}

	v := d.Canonical()
	if !bytes.Equal(v, []byte("https://example.com/")) {
		t.Fatalf("Invalid canonical URL: %s\n", v)
	}
}

func TestOpenGraph(t *testing.T) {

	d, err := html.ParseDocument([]byte(TestDocument))
	if err != nil {
		t.Fatalf("Failed to parse document: %s\n", err)
	}

	og := d.OpenGraph()

	t.Logf("\n%#v\n", *og)

}
