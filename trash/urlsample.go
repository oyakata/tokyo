package main

import (
	"fmt"
	"net/url"
)

func RelativeURL(base string, target string) {
	b, _ := url.Parse(base)
	u, _ := url.Parse(target)
	fmt.Printf("%v\n\t=> %v\n", base, b.ResolveReference(u))
}

func main() {
	var x *url.URL
	var y url.Values

	text := "https://www.google.co.jp/search?" +
		"q=%E4%BB%8A%E5%B7%9D%E7%BE%A9%E5%85%83%E3%81%AE%E3%82%BD%E3%83%8A%E3%82%BF" +
		"&ie=utf-8&oe=utf-8&client=firefox-b-ab" +
		"&gfe_rd=cr&ei=aqdLWJaKIabf8Afh3KKQDQ"

	x, _ = url.Parse(text)
	y, _ = url.ParseQuery(x.RawQuery)
	fmt.Println(y)

	RelativeURL("http://t.co/home/echizen/_go/src/github.com/oyakata/tokyo?x=1&y=2&z=9#ga",
				"../../../pa?ga=")
	RelativeURL("http://t.co/home/echizen/_go/src/github.com/oyakata/tokyo?x=1&y=2&z=9#ga",
				"//pa?ga=")
}
