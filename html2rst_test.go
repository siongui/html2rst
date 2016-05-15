package html2rst

import (
	iconv "github.com/djimenez/iconv-go"
	"net/http"
	"testing"
)

func TestHtmlToRst(t *testing.T) {
	//resp, err := http.Get("http://nanda.online-dhamma.net/")
	//resp, err := http.Get("https://siongui.github.io/zh/2016/03/14/pillow-useful-items-for-me-notes/")
	//resp, err := http.Get("http://myweb.ncku.edu.tw/~lsn46/Tipitaka/Sutta/Khuddaka/Khuddaka-patha/Khuddaka-patha.html")
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
	//print(HtmlToRst(resp.Body))
}
