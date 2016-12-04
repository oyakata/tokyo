package main

import (
	"math/rand"
	"./bench"
)

func makeParams(P int, Q int) (x, y []int) {
	var items []int
	var targets []int
	items = make([]int, P)
	for i:= 0; i < len(items); i++ {
		items[i] = rand.Intn(255)
	}

	targets = make([]int, rand.Intn(Q) + 1)
	for j:=0; j < len(targets); j++ {
		targets[j] = j
	}
	return items, targets
}

func main() {
	items, targets := makeParams(500000, 50000)
	bench.CountUp1(items, targets...)
}
