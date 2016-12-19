package main

import (
	"fmt"
)

func main() {
	// fmt.Println([]byte("𩸽のひらき"))
	// fmt.Println(string([]byte{240, 169, 184, 189, 227, 129, 174, 227, 129, 178, 227, 130, 137, 227, 129, 141}))
	fmt.Println(string([]byte{240, 169, 184, 189,
				  227, 129, 174,
				  227, 129, 178,
				  227, 130, 137,
				  227, 129, 141,
				  }))
}
