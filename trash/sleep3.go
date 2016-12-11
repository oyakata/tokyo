package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var acc, delay time.Duration
	rand.Seed(time.Now().UnixNano()) // ハマりポイント1
	choices := rand.Perm(9)

	for _, x := range choices[:5] {
		delay = time.Duration((x+1)*100) * time.Millisecond // ハマりポイント2
		acc += delay
		time.Sleep(delay)
	}
	fmt.Printf("合計: %v, 整数-> %v", acc, int(acc)+1000) // 合計: 2.8s, 整数-> 2800001000
}
