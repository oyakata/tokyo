package main;

import fmt "fmt"


func main() {
    xs := []int{1, 2, 3}

    for i:= 0; i < len(xs); i++ {
        fmt.Println(xs[i])
    }

    fmt.Println("====================")

    // 代入しないとダメ: append(xs, 9, 8, 7, 6, 5) evaluated but not used
    xs = append(xs, 9, 8, 7, 6, 5)

    for _, x := range xs {
        fmt.Println(x)
    }
}
