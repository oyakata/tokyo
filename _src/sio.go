package main

import (
	"fmt"
	"io/ioutil"
	"bytes"
	"bufio"
	"os"
)

var text = `
わたしは
い＊がわ
よ＊＊と
`

func main(){
	// 1. bytes.NewReaderを使う方法 (PythonのStringIOに相当)
	x := bufio.NewScanner(bytes.NewReader([]byte(text)))

	for x.Scan() {
		fmt.Println(x.Text())
	}

	// 2. ioutil.TempFileを使う方法 (PythonのNamedTemporaryFileに相当)
	tmp, _ := ioutil.TempFile("", "_")
	tmp.Write([]byte("| わたしは\n"))
	tmp.Write([]byte("| い＊がわ\n"))
	tmp.Write([]byte("| よ＊＊と"))
	defer os.Remove(tmp.Name()) // PythonのNamedTemporaryFileと違い、ファイルの削除は自分でやる。

	tmp.Seek(0, 0)
	x = bufio.NewScanner(tmp)

	for x.Scan() {
		fmt.Println(x.Text())
	}

	// 3. os.Openを使う方法 (Pythonのビルトインopen関数に相当)
	fp, _ := os.Open(tmp.Name())
	defer fp.Close()
	x = bufio.NewScanner(fp)
	for x.Scan() {
		fmt.Println(x.Text())
	}
}
