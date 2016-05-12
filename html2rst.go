package html2rst

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

func getAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func textNode2rst(n *html.Node) string {
	text := strings.TrimSpace(n.Data)
	if text == "" {
		return "\n"
	}
	return n.Data
}

func commentNode2rst(n *html.Node) string {
	rstText := "\n..\n"
	rstText += indentEachLine(strings.TrimSpace(n.Data))
	rstText += "\n"
	return rstText
}

func traverse(n *html.Node) string {
	if isTextNode(n) {
		return textNode2rst(n)
	}
	if isAnchorElement(n) {
		return a2rst(n)
	}
	if isImgElement(n) {
		return img2rst(n)
	}
	if isUlElement(n) {
		return ul2rst(n)
	}
	if isHrElement(n) {
		return hr2rst(n)
	}
	if isOlElement(n) {
		return ol2rst(n)
	}
	if isScriptElement(n) {
		return ""
	}
	if isCommentNode(n) {
		return commentNode2rst(n)
	}

	rstText := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		rstText += traverse(c)
	}

	return rstText
}

func HtmlToRst(r io.Reader) string {
	doc, err := html.Parse(r)
	if err != nil {
		panic("Fail to parse html")
	}

	return traverse(doc)
}
