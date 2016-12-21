package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math"
	"os"
	"os/exec"
)

func NewImage() *image.RGBA {
	var w, h int = 280, 240
	m := image.NewRGBA(image.Rect(0, 0, w, h))

	type Circle struct {
		X, Y, R float64
	}

	var Inside = func(c Circle, x, y int) bool {
		// sqrt((x-a)^2 + (y-b)^2) は二点(x, y), (a, b)間の距離
		// これが(x, y)を中心とする円の半径r以内の長さであれば円の内側

		xx, yy := c.X - float64(x), c.Y - float64(y)
		return math.Sqrt(xx*xx + yy*yy) <= c.R
	}

	large := Circle{140, 240, 80}
	middle := Circle{140, 180, 50}
	small := Circle{140, 120, 20}

	white := color.RGBA{
		255,
		255,
		255,
		255,
	}
	another := color.RGBA{
		170,
		179,
		0,
		255,
	}
	orange := color.RGBA{
		255,
		102,
		0,
		255,
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			var c color.RGBA
			if Inside(small, x, y) {
				// ちっちゃい円の中にある座標はオレンジ色(橙の絵)
				c = orange
			} else if Inside(large, x, y) || Inside(middle, x, y) {
				// 大きい円、または中くらいの円の中にある座標は白(お餅の絵)
				c = white
			} else {
				// それ以外は背景とする
				c = another
			}
			m.Set(x, y, c)
		}
	}
	return m
}

func main() {
	img := NewImage()

	tmp, err := ioutil.TempFile("", "SampleImage-")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer os.Remove(tmp.Name())
	png.Encode(tmp, img)
	exec.Command("eog", tmp.Name()).Run()
}
