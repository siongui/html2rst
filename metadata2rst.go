package html2rst

import (
	"golang.org/x/text/width"
	"io"
	"strings"
	"text/template"
)

var rstMetadataTmpl = `{{.Title}}
{{rstTitle .Title}}

:date: {{.Date}}
:modified: {{.Modified}}
:tags: {{.Tags}}
:category: {{.Category}}
:author: {{.Author}}
:summary: {{.Summary}}
:og_image: {{.OgImage}}
:oldurl: {{.OriginalURL}}


`

func GetWidthUTF8String(s string) int {
	size := 0
	for _, runeValue := range s {
		p := width.LookupRune(runeValue)
		if p.Kind() == width.EastAsianWide {
			size += 2
			continue
		}
		if p.Kind() == width.EastAsianNarrow {
			size += 1
			continue
		}
		if p.Kind() == width.EastAsianFullwidth {
			size += 2
			continue
		}
		if p.Kind() == width.EastAsianHalfwidth {
			size += 1
			continue
		}
		if p.Kind() == width.Neutral {
			size += 1
			continue
		}
		if p.Kind() == width.EastAsianAmbiguous {
			size += 1
			continue
		}
		panic("cannot determine!")
	}
	return size
}

func rstTitle(title string) string {
	width := GetWidthUTF8String(title)
	return strings.Repeat("=", width)
}

func writeRstMetadata(wr io.Writer, md *Metadata) {
	var funcMap = template.FuncMap{
		"rstTitle": rstTitle,
	}
	tmpl := template.Must(template.New("rstmeta").Funcs(funcMap).Parse(rstMetadataTmpl))
	tmpl.Execute(wr, md)
}
