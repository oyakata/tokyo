package main

import (
	"fmt"
	"io/ioutil"
	"io"
	"bytes"
	"bufio"
	"os"
)

var text = `わたしは
い＊がわ
よ＊＊と
`

type CustomScanner struct {
	s *bufio.Scanner
}

func (c *CustomScanner) PrintAll() {
	number := 1
	for c.s.Scan() {
		fmt.Printf("%v:\t%v\n", number, c.s.Text())
		number++
	}
	fmt.Println("--------------------")
}

func NewCustomScanner(rd *io.Reader) *CustomScanner {
	c := CustomScanner{}
	c.s = bufio.NewScanner(*rd)
	return &c
}

func main(){
	var rd io.Reader
	var x *CustomScanner

	rd = bytes.NewReader([]byte(text))
	x = NewCustomScanner(&rd)

	x.PrintAll()

	tmp, _ := ioutil.TempFile("", "_")
	tmp.Write([]byte("| わたしは\n"))
	tmp.Write([]byte("| い＊がわ\n"))
	tmp.Write([]byte("| よ＊＊と"))
	defer os.Remove(tmp.Name())

	tmp.Seek(0, 0)
	rd = tmp
	x = NewCustomScanner(&rd)

	x.PrintAll()

	fp, _ := os.Open(tmp.Name())
	defer fp.Close()
	rd = fp
	x = NewCustomScanner(&rd)
	x.PrintAll()
}
