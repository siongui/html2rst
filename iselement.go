package html2rst

import "golang.org/x/net/html"

func isAnchorElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

func isUlElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "ul"
}

func isLiElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "li"
}

func isScriptElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "script"
}

func isImgElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "img"
}

func isHrElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "hr"
}

func isOlElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "ol"
}

func isTextNode(n *html.Node) bool {
	return n.Type == html.TextNode
}

func isCommentNode(n *html.Node) bool {
	return n.Type == html.CommentNode
}
