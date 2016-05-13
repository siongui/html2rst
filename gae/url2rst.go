package gaehtml2rst

import (
	"net/http"

	"github.com/siongui/html2rst"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func url2Rst(url string, r *http.Request) string {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return html2rst.HtmlToRst(resp.Body)
}
