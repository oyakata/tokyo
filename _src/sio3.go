package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"bufio"
	"bytes"
	"os"
)

func main(){
	var reader io.Reader
	var scanner *bufio.Scanner

	reader = bytes.NewReader([]byte("Hello, world."))
	scanner = bufio.NewScanner(reader)
	fmt.Println(scanner)

	tmp, _ := ioutil.TempFile("", "_")
	defer os.Remove(tmp.Name())
	tmp.Write([]byte("I have a pen."))
	tmp.Seek(0, 0)
	
	reader = tmp
	scanner = bufio.NewScanner(reader)
	fmt.Println(scanner)

	fp, _ := os.Open(tmp.Name())
	defer fp.Close()
	reader = fp
	scanner = bufio.NewScanner(reader)
	fmt.Println(scanner)
}
