package bench

import (
	"math/rand"
	"testing"
)

func makeParams(P int, Q int) (x, y []int) {
	var items []int
	var targets []int
	items = make([]int, P)
	for i := 0; i < len(items); i++ {
		items[i] = rand.Intn(Q)
	}

	targets = make([]int, Q)
	for j := 0; j < len(targets); j++ {
		targets[j] = j
	}
	return items, targets
}

const ITEMS1 = 50000
const TARGETS1 = 50

const ITEMS2 = 50000
const TARGETS2 = 50000

const ITEMS3 = 50
const TARGETS3 = 50000

func BenchmarkTest1Using1(b *testing.B) {
	items, targets := makeParams(ITEMS1, TARGETS1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountUp1(items, targets...)
	}
}

func BenchmarkTest1Using2(b *testing.B) {
	items, targets := makeParams(ITEMS1, TARGETS1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountUp2(items, targets...)
	}
}

func BenchmarkTest2Using1(b *testing.B) {
	items, targets := makeParams(ITEMS2, TARGETS2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountUp1(items, targets...)
	}
}

func BenchmarkTest2Using2(b *testing.B) {
	items, targets := makeParams(ITEMS2, TARGETS2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountUp2(items, targets...)
	}
}

func BenchmarkTest3Using1(b *testing.B) {
	items, targets := makeParams(ITEMS3, TARGETS3)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountUp1(items, targets...)
	}
}

func BenchmarkTest3Using2(b *testing.B) {
	items, targets := makeParams(ITEMS3, TARGETS3)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountUp2(items, targets...)
	}
}
