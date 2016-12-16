package main

import "fmt"

func main() {
	var hokke, zok rune

	hokke, zok = '\U00029E3D', '\U0002023B'
	fmt.Printf("%c, %c", hokke, zok) // 𩸽, 𠈻
}
