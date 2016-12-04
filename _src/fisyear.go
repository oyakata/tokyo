package main

import fmt "fmt"

// この相対パスのインポートはダメだと思う。
//	=> GOPATH, GOROOTの使い方がまだわからない。
import api "./tyo/datetime"

func main() {
	value := api.FiscalYear(2016)
	fmt.Printf("Hello, world[%v].\n", value)
	fmt.Println(fmt.Sprintf("Hello, world[%v].\n", value))
}
