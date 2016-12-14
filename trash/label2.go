package main

import (
	"fmt"
	"strings"
	"time"
)

var ora = []string{
	"ハァ テレビもねぇ！ ラジオもねぇ！",
	"クルマもそれほど走ってねぇ！",
	"ピアノもねぇ！ バーもねぇ！ おまわり毎日ぐーるぐる！",
}

const delay = time.Duration(750 * time.Millisecond)

func Sing(message string) {
	xs := strings.Split(message, " ")
	length := len(xs)
	for i, x := range xs {
		fmt.Print(x)
		if i+1 < length {
			fmt.Print(" ")
			time.Sleep(delay)
		}
	}
	time.Sleep(delay)
	fmt.Println()
}

func main() {

	for i := 0; true; i++ {

		Sing(ora[i])

		if i >= 2 {
			Sing("おらこんな村いやだぁ〜 おらこんな村いやだぁ〜")
			Sing("東京へ出るだぁ〜〜〜")
			goto Tokyo
		}
	}

Tokyo:

	for i := 0; true; i++ {
		Sing("東京へ出だな〜ら 銭コア貯めでぇ 東京でベコ飼うだぇ〜〜")
		break Tokyo
	}
}
