package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	text := `いまがわ
	よしもと
	にっぽんいち
	`
	bs := bufio.NewScanner(bytes.NewReader([]byte(text)))
	ss := bufio.NewScanner(strings.NewReader(text))

	for bs.Scan() {
		fmt.Println("bs:", bs.Text())
	}

	for ss.Scan() {
		fmt.Println("ss", ss.Text())
	}

	// bs: いまがわ
	// bs: 	よしもと
	// bs: 	にっぽんいち
	// bs:
	// ss いまがわ
	// ss 	よしもと
	// ss 	にっぽんいち
	// ss
}
