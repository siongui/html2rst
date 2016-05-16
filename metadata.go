package html2rst

import (
	"bytes"
	"golang.org/x/net/html"
	"io"
	"strings"
)

type Metadata struct {
	Title       string
	Date        string
	Modified    string
	Tags        string
	Category    string
	Author      string
	Summary     string
	OriginalURL string
	OgImage     string
}

func traverse4metadata(n *html.Node, md *Metadata) {
	if isTitleElement(n) {
		md.Title = n.FirstChild.Data
	}
	if isMetaElement(n) {
		name, ok := getAttribute(n, "name")
		if ok {
			if strings.ToLower(name) == "keywords" {
				md.Tags, _ = getAttribute(n, "content")
			}
			if strings.ToLower(name) == "description" {
				md.Summary, _ = getAttribute(n, "content")
			}
			if strings.ToLower(name) == "author" {
				md.Author, _ = getAttribute(n, "content")
			}
		}
		property, ok := getAttribute(n, "property")
		if ok {
			if property == "og:image" {
				md.OgImage, _ = getAttribute(n, "content")
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse4metadata(c, md)
	}
}

func GetHtmlMetadata(r io.Reader, url string) string {
	doc, err := html.Parse(r)
	if err != nil {
		panic("Fail to parse html")
	}

	md := Metadata{}
	md.OriginalURL = url

	traverse4metadata(doc, &md)

	var rstmeta bytes.Buffer
	writeRstMetadata(&rstmeta, &md)
	return rstmeta.String()
}
