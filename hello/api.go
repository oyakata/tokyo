package hello

import "fmt"

func GotoTokyo() {
	for _, x := range [7]int{1, 2, 3, 4, 5, 6, 7} {
		fmt.Printf("go> to %v\n", x)
	}
}
