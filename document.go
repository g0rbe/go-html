package html

import (
	"bytes"
	"io"

	"github.com/PuerkitoBio/goquery"
)

type Document struct {
	s *goquery.Selection
}

func ReadDocument(r io.Reader) (*Document, error) {

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	d := new(Document)

	d.s = doc.Unwrap()

	return d, nil
}

func ParseDocument(b []byte) (*Document, error) {
	buf := bytes.NewReader(b)
	return ReadDocument(buf)
}

// Text returns the first text element that match selector.
func (d *Document) Text(selector string) []byte {
	v := d.s.FindMatcher(goquery.Single(selector)).Text()
	return []byte(v)
}

// Attribute returns the content of attribute from the first element that match selector.
//
// If attribute not exist, returns nil.
func (d *Document) Attribute(selector, attribute string) []byte {
	v, exist := d.s.FindMatcher(goquery.Single(selector)).Attr(attribute)
	if !exist {
		return nil
	}
	return []byte(v)
}

// InnerHTML returns the inner HTML of the first element that match selector.
//
// If not found or cant render the HTML, returns nil.
func (d *Document) InnerHTML(selector string) []byte {
	v, err := goquery.OuterHtml(d.s.FindMatcher(goquery.Single(selector)))
	if err != nil {
		return nil
	}
	return []byte(v)
}

// OuterHTML returns the outer HTML of the first element that match selector.
//
// If not found or cant render the HTML, returns nil.
func (d *Document) OuterHTML(selector string) []byte {
	v, err := goquery.OuterHtml(d.s.FindMatcher(goquery.Single(selector)))
	if err != nil {
		return nil
	}
	return []byte(v)
}

// RemoveSelection removes elements from d Document that matches the selector.
func (d *Document) RemoveSelection(selector string) {
	d.s.Find(selector).Remove()
}

func (d *Document) Title() []byte {
	return d.Text(TitleSelector)
}

func (d *Document) Description() []byte {
	return d.Attribute(DescriptionSelector, "content")
}

func (d *Document) Canonical() []byte {
	return d.Attribute(CanonicalSelector, "href")
}

func (d *Document) OpenGraph() *OpenGraph {

	if d.s.Find(OpenGraphSelector).Length() == 0 {
		return nil
	}

	og := new(OpenGraph)

	og.URL = string(d.Attribute(OpenGraphURLSelector, "content"))
	og.SiteName = string(d.Attribute(OpenGraphSiteNameSelector, "content"))
	og.Title = string(d.Attribute(OpenGraphTitleSelector, "content"))
	og.Description = string(d.Attribute(OpenGraphDescriptionSelector, "content"))
	og.Locale = string(d.Attribute(OpenGraphLocaleSelector, "content"))
	og.Type = string(d.Attribute(OpenGraphTypeSelector, "content"))
	og.Image = string(d.Attribute(OpenGraphImageSelector, "content"))

	return og
}
