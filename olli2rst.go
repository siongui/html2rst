package html2rst

import (
	"golang.org/x/net/html"
	"strconv"
)

func liInOl2rst(n *html.Node, orderCount int) string {
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

	return strconv.Itoa(orderCount) + ". " + rstText + "\n"
}

func ol2rst(n *html.Node) string {
	rstText := ""
	orderCount := 1
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if isLiElement(c) {
			rstText += liInOl2rst(c, orderCount)
			orderCount++
		}
	}

	return rstText
}
