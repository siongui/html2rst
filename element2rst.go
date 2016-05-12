package html2rst

import (
	"golang.org/x/net/html"
	"strings"
)

func a2rst(n *html.Node) string {
	if isImgElement(n.FirstChild) {
		rstText := img2rst(n.FirstChild)
		href, ok := getAttribute(n, "href")
		if ok {
			rstText += "   :target: "
			rstText += href
			rstText += "\n"
		}
		return rstText
	}

	text := strings.TrimSpace(n.FirstChild.Data)

	href, ok := getAttribute(n, "href")
	if ok {
		return "`" + text + " <" + href + ">`__"
	}

	return ""
}

func img2rst(n *html.Node) string {
	rstText := ".. image:: "

	src, ok := getAttribute(n, "src")
	if ok {
		rstText += src
		rstText += "\n"
	} else {
		rstText += "\n"
	}

	alt, ok := getAttribute(n, "alt")
	if ok {
		rstText += "   :alt: "
		rstText += alt
		rstText += "\n"
	}

	class, ok := getAttribute(n, "class")
	if ok {
		if class == "align-center" {
			rstText += "   :align: center\n"
		}
	}

	return rstText
}

func hr2rst(n *html.Node) string {
	return "\n----\n"
}
