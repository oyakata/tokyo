package bench

type Pair [2]int

// itemsの中からtargetに該当するものを数える
func CountUp1(items []int, targets ...int) []Pair {
	result := make([]Pair, 0, len(targets))

	var acc int
	for _, x := range targets {
		acc = 0
		for _, y := range items {
			if x == y {
				acc++
			}
		}
		result = append(result, Pair{x, acc})
	}
	return result
}

func CountUp2(items []int, targets ...int) []Pair {
	xs := make(map[int]int, len(items))
	for _, x := range items {
		xs[x] += 1
	}

	result := make([]Pair, 0, len(targets))
	for _, y := range targets {
		result = append(result, Pair{y, xs[y]})
	}
	return result
}

//
// func main() {
// 	items := []int{9, 10, 10, 1, 10, 8}
// 	targets := []int{10, 1}
// 	fmt.Println(CountUp1(items, targets...))
// 	fmt.Println(CountUp2(items, targets...))
// }
