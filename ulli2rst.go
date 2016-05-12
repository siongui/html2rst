package html2rst

import (
	"golang.org/x/net/html"
)

func li2rst(n *html.Node) string {
	rstText := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if isTextNode(c) {
			rstText += textNode2rst(c)
		}
		if isAnchorElement(c) {
			rstText += a2rst(c)
		}
		if isUlElement(c) {
			rstText += "\n"
			rstText += indentEachLine(ul2rst(c))
		}
	}

	return "- " + rstText + "\n"
}

func ul2rst(n *html.Node) string {
	rstText := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if isLiElement(c) {
			rstText += li2rst(c)
		}
	}

	return rstText
}
