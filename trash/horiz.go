package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"os/exec"
)

func NewImage() *image.RGBA {
	var w, h int = 280, 240
	m := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {

			var value uint8 = 0
			// 10pxの底辺1本, 2pxの線1本
			if y > 220 || y >= 180 && y <= 182 {
				value = 255
			}

			c := color.RGBA{
				value,
				value,
				value,
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
