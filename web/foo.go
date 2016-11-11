package main

// htmlパッケージが外部パッケージなのでgo getで入れないと動かない

// GOPATH=`pwd`/.go go get
// GOPATH=`pwd`/.go go run htm.go
// GOPATHは絶対パスしか受け付けてくれない。相対パスを渡すとエラーになる
import (
	"fmt"
	"golang.org/x/net/html"
	"bytes"
)


var text = `
<html>
<body>
This is here document.

ヒアドキュメント的な
複数行文字列作成方法
（Raw文字列でもある）
</body>
</html>
`


func main(){
	var fp bytes.Buffer
	fp.WriteString(text)
	// fmt.Println(fp.String())
	// rs := bytes.NewReader(fp.Bytes())
	rs := bytes.NewReader([]byte(text))
	doc, err := html.Parse(rs)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(doc)
}
