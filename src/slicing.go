package main

import (
	"fmt"
)

func main() {
	L := make([]string, 3, 5)
	L[0], L[1], L[2] = "今川", "義元", "スルガ"
	fmt.Printf("len=%v, cap=%v: %v\n", len(L), cap(L), L)
	LL := append(L, "治部大輔", "海道", "一の", "弓取り")
	fmt.Printf("len=%v, cap=%v: %v\n", len(LL), cap(LL), LL)
	fmt.Printf("??> len=%v, cap=%v: %v\n", len(L), cap(L), L)
	/*
	len=3, cap=5: [今川 義元 スルガ]
	len=7, cap=10: [今川 義元 スルガ 治部大輔 海道 一の 弓取り]
	??> len=3, cap=5: [今川 義元 スルガ]
	*/
}

