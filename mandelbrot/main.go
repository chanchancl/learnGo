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
	centerX                = 0
	centerY                = 0
	size                   = 2
	xmin, ymin, xmax, ymax = centerX - size, centerY - size, centerX + size, centerY + size
	width, height          = 5024, 5024
)

var wg sync.WaitGroup

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))

    ngo := 20
    for i :=0; i < ngo; i++ {
        dx := width / ngo
        go getPic(img.SubImage(image.Rect( dx * i, 0, dx * (i+1), height)),dx * i, 0, dx * (i+1), height)
        wg.Add(1)
    }
    
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
	const iterations = 256
	const contrast = 15

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return getColor(n)
		}
	}
	return color.Black
}

func getColor(n int) color.Color {
    paletted := [16]color.Color{
        color.RGBA{66, 30, 15, 255},    // # brown 3
        color.RGBA{25, 7, 26, 255},     // # dark violett
        color.RGBA{9, 1, 47, 255},      // # darkest blue
        color.RGBA{4, 4, 73, 255},      // # blue 5
        color.RGBA{0, 7, 100, 255},     // # blue 4
        color.RGBA{12, 44, 138, 255},   // # blue 3
        color.RGBA{24, 82, 177, 255},   // # blue 2
        color.RGBA{57, 125, 209, 255},  // # blue 1
        color.RGBA{134, 181, 229, 255}, // # blue 0
        color.RGBA{211, 236, 248, 255}, // # lightest blue
        color.RGBA{241, 233, 191, 255}, // # lightest yellow
        color.RGBA{248, 201, 95, 255},  // # light yellow
        color.RGBA{255, 170, 0, 255},   // # dirty yellow
        color.RGBA{204, 128, 0, 255},   // # brown 0
        color.RGBA{153, 87, 0, 255},    // # brown 1
        color.RGBA{106, 52, 3, 255},    // # brown 2
    }
    return paletted[n%16]
}
