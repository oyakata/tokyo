package main

import (
	"fmt"
	"time"
)

func main() {
	var x time.Duration
	x = 500
	time.Sleep(x * time.Millisecond)
	fmt.Printf("Waited: %v", x*time.Millisecond) // Waited: 500ms
}
