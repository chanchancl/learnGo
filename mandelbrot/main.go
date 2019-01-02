package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

const (
	centerX                = -0.540046768
	centerY                = 0.540046768
	size                   = 0.0010000
	xmin, ymin, xmax, ymax = centerX - size, centerY - size, centerX + size, centerY + size
	width, height          = 1024, 1024
)

var wg sync.WaitGroup

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	wg.Add(4)
	go getPic(img.SubImage(image.Rect(0, 0, width/2, height/2)), 0, 0, width/2, height/2)
	go getPic(img.SubImage(image.Rect(width/2, 0, width, height/2)), width/2, 0, width, height/2)
	go getPic(img.SubImage(image.Rect(0, height/2, width/2, height)), 0, height/2, width/2, height)
	go getPic(img.SubImage(image.Rect(width/2, height/2, width, height)), width/2, height/2, width, height)
	wg.Wait()

	png.Encode(os.Stdout, img)
}

func getPic(pimg image.Image, left, bottom, right, top int) {
	//localHeight := top - bottom
	//localWidth := right - left

	img := pimg.(*image.RGBA)
	for py := bottom; py < top; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := left; px < right; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	wg.Done()
}

func mandelbrot(z complex128) color.Color {
	const iterations = 255
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
