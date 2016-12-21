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

// func NewImage() *image.RGBA {
// 	img := image.NewRGBA(image.Rect(0, 0, 100, 50))
// 	img.Set(2, 3, color.RGBA{255, 0, 0, 255})
// 	return img
// }

type Circle struct {
	X, Y, R float64
}

func (c *Circle) Brightness(x, y float64) uint8 {
	// 点 (a, b) を中心とする半径 r の円は
	// (x - a)^2 + (y - b)^2 = r^2

	// math.Sqrt(dx*dx+dy*dy) は二点間の距離

	var dx, dy float64 = c.X - x, c.Y - y
	d := math.Sqrt(dx*dx+dy*dy) / c.R
	if d > 1 {
		return 0
	} else {
		return 255
	}

	// if dx*dx + dy*dy <= c.R * c.R {
	// 	return 255
	// } else {
	// 	return 0
	// }
}

func NewImage() *image.RGBA {
	// キャンバスの幅と高さ
	var w, h int = 280, 240

	// 原点
	var hw, hh float64 = float64(w / 2), float64(h / 2)

	// 半径が40
	r := 40.0
	// r := 60.0

	// θ := 2 * math.Pi / 3 // 360° = 2π  => 2/3π = 120°のラジアン
	// θ := 120 * math.Pi / 180

	// cr := &Circle{hw - r*math.Sin(0), hh - r*math.Cos(0), 60}
	// cg := &Circle{hw - r*math.Sin(θ), hh - r*math.Cos(θ), 60}
	// cb := &Circle{hw - r*math.Sin(-θ), hh - r*math.Cos(-θ), 60}

	var short = func(A, B, C int) (*Circle, *Circle, *Circle) {
		one, two, three := float64(A), float64(B), float64(C)
		cr := &Circle{hw - r*math.Cos(one * math.Pi / 180), hh - r*math.Sin(one * math.Pi / 180), 60}
		cg := &Circle{hw - r*math.Cos(two * math.Pi / 180), hh - r*math.Sin(two * math.Pi / 180), 60}
		cb := &Circle{hw - r*math.Cos(three * math.Pi / 180), hh - r*math.Sin(three * math.Pi / 180), 60}
		return cr, cg, cb
	}

	cr, cg, cb := short(90, 330, 210)

	m := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := color.RGBA{
				cr.Brightness(float64(x), float64(y)),
				cg.Brightness(float64(x), float64(y)),
				cb.Brightness(float64(x), float64(y)),
				255,
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
