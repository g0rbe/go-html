package html

import "fmt"

// RemoveSelections removes elements from HTML document doc that matches any of the selectors.
func RemoveSelections(doc []byte, selectors []string) ([]byte, error) {

	d, err := ParseDocument(doc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse: %w", err)
	}

	d.RemoveSelections(selectors)

	r := d.HTML()
	if r == nil {
		return nil, fmt.Errorf("cant render HTML")
	}

	return r, nil
}
