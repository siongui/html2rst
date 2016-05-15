package html2rst

import (
	"golang.org/x/net/html"
)

func td2rst(td *html.Node) string {
	s := ""
	for c := td.FirstChild; c != nil; c = c.NextSibling {
		if isTextNode(c) {
			s += textNode2rst(c)
			continue
		}
		if isAnchorElement(c) {
			s += a2rst(c)
			continue
		}
	}
	return s
}

func tr2rst(tr *html.Node) string {
	s := ""
	isFirstTd := true
	for c := tr.FirstChild; c != nil; c = c.NextSibling {
		if isTdElement(c) {
			if isFirstTd {
				s += ("  * - " + td2rst(c) + "\n")
				isFirstTd = false
			} else {
				s += ("    - " + td2rst(c) + "\n")
			}
			continue
		}
	}
	return s
}

func tbody2rst(tbody *html.Node) string {
	s := ""
	for c := tbody.FirstChild; c != nil; c = c.NextSibling {
		if isTrElement(c) {
			s += tr2rst(c)
			continue
		}
	}
	return s
}

func table2rst(table *html.Node) string {
	s := ".. list-table::\n\n"
	for c := table.FirstChild; c != nil; c = c.NextSibling {
		if isTbodyElement(c) {
			s += tbody2rst(c)
			continue
		}
	}
	return s
}
