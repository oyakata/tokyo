package main

import (
	ikzo "github.com/oyakata/ikzo/hello"
	"github.com/oyakata/tokyo/good"
	tyo "github.com/oyakata/tokyo/hello"
	"github.com/oyakata/tokyo/gdbsample"
)

func main() {
	ikzo.World()
	tyo.World()
	tyo.GotoTokyo()
	// good.DoBest(1000)  // パッケージ宣言は best なので、goodという名前では参照できない
	// # command-line-arguments
	// ./main_tyo.go:6: imported and not used: "github.com/oyakata/tokyo/good" as best
	// ./main_tyo.go:13: undefined: good in good.DoBest
	best.DoBest(1000)
	gdbsample.Sample()
}
