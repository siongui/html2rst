package html2rst

import (
	iconv "github.com/djimenez/iconv-go"
	"net/http"
	"testing"
)

func TestHtmlToRst(t *testing.T) {
	resp, err := http.Get("http://myweb.ncku.edu.tw/~lsn46/Tipitaka/Sutta/Digha/Saamannaphala/Saamannaphala.html")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	print(HtmlToRst(resp.Body))
}

func TestNonUtf8HtmlToRst(t *testing.T) {
	return
	resp, err := http.Get("http://myweb.ncku.edu.tw/~lsn46/Tipitaka/Samyutta-nikaaya.htm")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reader, err := iconv.NewReader(resp.Body, "big5", "utf-8")
	if err != nil {
		panic(err)
	}
	print(HtmlToRst(reader))
}
