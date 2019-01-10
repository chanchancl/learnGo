package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const (
	Width        = 512
	Height       = 512
	N            = 128
	MAX_STEP     = 10
	MAX_DISTANCE = 2.0
	EPSILON      = 1e-6
)

func circleSDF(x, y, cx, cy, r float64) float64 {
	ux, uy := x-cx, y-cy
	return math.Sqrt(ux*ux+uy*uy) - r
}

func trace(ox, oy, dx, dy float64) float64 {
	var t float64
	t = 0
	for i := 0; i < MAX_STEP && t < MAX_DISTANCE; i++ {
		sd := circleSDF(ox+dx*t, oy+dy*t, 0.5, 0.5, 0.1)
		if sd < EPSILON {
			return 2.0
		}
		t += sd
	}
	return 0.0
}

func sample(x, y float64) float64 {
	sum := 0.0
	for i := 0; i < N; i++ {
		//a := 2 * math.Pi * rand.Float64()
		//a := 2 * math.Pi * float64(i) / N
		a := 2 * math.Pi * (float64(i) + rand.Float64()) / N
		sum += trace(x, y, math.Cos(a), math.Sin(a))
	}
	return sum / N
}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))

	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			value := (uint8)(math.Min(sample(float64(x)/Width, float64(y)/Height)*255.0, 255.0))
			img.Set(x, y, color.Gray{value})
		}
	}

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dirname := filepath.Base(dir)
	dirpath, _ := filepath.Abs(dir + "/..")

	timestr := time.Now().Format("2006-01-02 15-04-05")
	filename := fmt.Sprintf("%s-%s.png", dirname, timestr)

	path := filepath.Join(dirpath, filename)
	fmt.Println(path)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	png.Encode(file, img)
	file.Close()
}
