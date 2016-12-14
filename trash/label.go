package main

import (
	"fmt"
	"time"
)

func main() {
L:
	for i := 1; true; i++ {

		for j := 1; true; j++ {

			fmt.Printf("i=%v> j=%v\n", i, j)

			if i == 2 && j == 2 {
				continue L
			}

			if i%2 != 0 && j == 3 {
				goto L2
			}
			if j == 3 {
				break L
			}

			time.Sleep(200 * time.Millisecond)
		}

	L2:
	}
}
