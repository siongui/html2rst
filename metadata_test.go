package html2rst

import (
	"net/http"
	"testing"
)

func TestGetHtmlMetadata(t *testing.T) {
	url := "http://myweb.ncku.edu.tw/~lsn46/Tipitaka/Sutta/Digha/Saamannaphala/Saamannaphala.html"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	println(GetHtmlMetadata(resp.Body, url))
}
