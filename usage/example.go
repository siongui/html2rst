package main

import (
	"github.com/siongui/html2rst"
	"net/http"
)

func main() {
	//resp, err := http.Get("http://nanda.online-dhamma.net/")
	//resp, err := http.Get("https://siongui.github.io/zh/2016/03/14/pillow-useful-items-for-me-notes/")
	resp, err := http.Get("http://myweb.ncku.edu.tw/~lsn46/Tipitaka/Sutta/Khuddaka/Khuddaka-patha/Khuddaka-patha.html")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	print(html2rst.HtmlToRst(resp.Body))
}
